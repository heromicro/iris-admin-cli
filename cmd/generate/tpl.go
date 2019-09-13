package generate


// TplItem 模板项
type TplItem struct {
	StructName string         `json:"struct_name"` // 结构体名称
	Comment    string         `json:"comment"`     // 注释
	Fields     []TplFieldItem `json:"fields"`      // 字段项
}

func (t TplItem) toSchemaFields() []schemaField {
	var items []schemaField
	for _, f := range t.Fields {
		items = append(items, schemaField{
			Name:       f.StructFieldName,
			Comment:    f.Comment,
			Type:       f.StructFieldType,
			IsRequired: f.StructFieldRequired,
		})
	}
	return items
}

func (t TplItem) toEntityFields() []entityField {
	var items []entityField
	for _, f := range t.Fields {
		items = append(items, entityField{
			Name:        f.StructFieldName,
			Comment:     f.Comment,
			Type:        f.StructFieldType,
			GormOptions: f.GormOptions,
		})
	}
	return items
}

// TplFieldItem 模板字段项
type TplFieldItem struct {
	StructFieldName     string `json:"struct_field_name"`     // 结构体字段名称
	StructFieldRequired bool   `json:"struct_field_required"` // 结构字段必选项
	Comment             string `json:"comment"`               // 注释
	StructFieldType     string `json:"struct_field_type"`     // 结构体字段类型
	GormOptions         string `json:"gorm_options"`          // gorm配置项
}
