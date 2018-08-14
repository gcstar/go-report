package service

import (
	. "report/goreport/backend/db"
	"strings"
)

const SEPERATOR = "/"

type Widget struct {
	WidgetId     uint   `json:"widget_id" gorm:"primary_key"`
	UserId       string `json:"userId"`
	CategoryName string `json:"categoryName"`
	WidgetName   string `json:"widgetName"`
	DataJson     string `json:"dataJson"`
}

type CategoryName struct {
	CategoryName string
	WidgetName   string
}

type CategoryTreeNode struct {
	Label    string             `json:"label"`
	Children []CategoryTreeNode `json:"children"`
	Leaf     bool               `json:"leaf"`
}

func (Widget) TableName() string {
	return "dashboard_widget"
}

func FindWidget(fieldName string, keyWord string) *Widget {
	var widget Widget
	DB.Where(fieldName+"=?", keyWord).Find(&widget)
	return &widget
}

func GetCategoryWidgets() []CategoryTreeNode {
	var categoryNames []CategoryName
	DB.Raw("select distinct category_name as category_name from dashboard_widget").Scan(&categoryNames)

	var categoryTreeNodes []CategoryTreeNode
	for _, category := range categoryNames {
		if !strings.Contains(category.CategoryName, SEPERATOR) {
			categoryNode := CategoryTreeNode{Label: category.CategoryName, Leaf: false}
			categoryTreeNodes = append(categoryTreeNodes, categoryNode)
		}
	}
	for _, category := range categoryNames {
		if strings.Contains(category.CategoryName, SEPERATOR) {
			names := strings.Split(category.CategoryName, SEPERATOR)
			var parentName = names[0]
			var subName = names[1:]
			categoryTreeNodes = appendCategory(parentName, subName, categoryTreeNodes)
		}
	}
	DB.Raw("select category_name,widget_name  from dashboard_widget").Scan(&categoryNames)
	addWidgetToCategory(categoryNames, categoryTreeNodes)
	return categoryTreeNodes
}

func addWidgetToCategory(categorys []CategoryName, nodes []CategoryTreeNode) {
	var len = len(nodes)
	for _, category := range categorys {
		name := category.CategoryName
		for j := 0; j < len; j++ {
			if name == nodes[j].Label {
				nodes[j].Children = append(nodes[j].Children, CategoryTreeNode{Leaf: true, Label: category.WidgetName})
			} else {
				categoryNames := strings.Split(name, SEPERATOR)
				parentName := categoryNames[0]
				subNames := categoryNames[1:]
				if parentName == nodes[j].Label {
					addChildrenWidgets(category, subNames, nodes[j].Children)
				}
			}
		}
	}

}

func addChildrenWidgets(category CategoryName, subNames []string, children []CategoryTreeNode) {
	if len(subNames) == 0 {
		return
	}
	var parentName = subNames[0]
	for j := 0; j < len(children); j++ {
		if parentName == children[j].Label && len(subNames[1:]) == 0 {
			children[j].Children = append(children[j].Children, CategoryTreeNode{Label: category.WidgetName, Leaf: true})
		} else {
			addChildrenWidgets(category, subNames[1:], children[j].Children)
		}
	}
}

func appendCategory(parentName string, subName []string, nodes []CategoryTreeNode) []CategoryTreeNode {
	var flag = false
	var nodeLen = len(nodes)
	for j := 0; j < nodeLen; j++ {
		if parentName == nodes[j].Label {
			flag = true
		}
	}
	if !flag {
		var node = CategoryTreeNode{Label: parentName}
		nodes = append(nodes, node)
	}
	nodeLen = len(nodes)
	for j := 0; j < nodeLen; j++ {
		if parentName == nodes[j].Label {
			nodes[j].Children = append(nodes[j].Children, CategoryTreeNode{Label: subName[0], Leaf: false})
			if cap(subName[1:]) != 0 {
				appendCategory(subName[0], subName[1:], nodes[j].Children)
			}
		}
	}

	return nodes
}
