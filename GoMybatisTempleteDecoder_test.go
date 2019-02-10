package GoMybatis

import (
	"fmt"
	"testing"
)

func TestGoMybatisTempleteDecoder_Decode(t *testing.T) {
	var decoder = GoMybatisTempleteDecoder{}
	var mapElements = make([]ElementItem, 0)
	mapElements = append(mapElements, ElementItem{

	})
	var BaseResultMap = MapperXml{
		Tag: "resultMap",
		Id:  "BaseResultMap",
		ElementItems: []ElementItem{
			{
				ElementType: "id",
				Propertys: map[string]string{
					"column":   "id",
					"property": "id",
				},
			},
			{
				ElementType: "result",
				Propertys: map[string]string{
					"column":   "name",
					"property": "name",
					"langType": "string",
				},
			},
			{
				ElementType: "result",
				Propertys: map[string]string{
					"column":   "pc_link",
					"property": "pcLink",
					"langType": "string",
				},
			},
			{
				ElementType: "result",
				Propertys: map[string]string{
					"column":   "h5_link",
					"property": "h5Link",
					"langType": "string",
				},
			},
			{
				ElementType: "result",
				Propertys: map[string]string{
					"column":   "remark",
					"property": "remark",
					"langType": "string",
				},
			},
			{
				ElementType: "result",
				Propertys: map[string]string{
					"column":         "version",
					"property":       "version",
					"langType":       "int",
					"enable_version": "true",
				},
			},
			{
				ElementType: "result",
				Propertys: map[string]string{
					"column":   "create_time",
					"property": "createTime",
					"langType": "time.Time",
				},
			},
			{
				ElementType: "result",
				Propertys: map[string]string{
					"column":               "delete_flag",
					"property":             "deleteFlag",
					"langType":             "int",
					"enable_version":       "true",
					"enable_logic_delete":  "true",
					"logic_deleted_value":  "1",
					"logic_undelete_value": "0",
				},
			},
		},
	}
	fmt.Println(BaseResultMap)

	var xml = MapperXml{
		Tag: "selectTemplete",
		Propertys: map[string]string{
			"table":   "biz_activity",
			"columns": "*",
			"wheres":  "name?name = #{name}",
		},
		ElementItems: []ElementItem{},
	}
	var e = decoder.Decode(&xml)
	if e != nil {
		t.Fatal(e)
	}
}