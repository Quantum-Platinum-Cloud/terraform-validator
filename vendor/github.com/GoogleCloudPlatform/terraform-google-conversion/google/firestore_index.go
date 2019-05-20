// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import "reflect"

func GetFirestoreIndexCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//firestore.googleapis.com/{{name}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetFirestoreIndexApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "firestore.googleapis.com/Index",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/firestore/v1/rest",
				DiscoveryName:        "Index",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetFirestoreIndexApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	databaseProp, err := expandFirestoreIndexDatabase(d.Get("database"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("database"); !isEmptyValue(reflect.ValueOf(databaseProp)) && (ok || !reflect.DeepEqual(v, databaseProp)) {
		obj["database"] = databaseProp
	}
	collectionProp, err := expandFirestoreIndexCollection(d.Get("collection"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("collection"); !isEmptyValue(reflect.ValueOf(collectionProp)) && (ok || !reflect.DeepEqual(v, collectionProp)) {
		obj["collection"] = collectionProp
	}
	queryScopeProp, err := expandFirestoreIndexQueryScope(d.Get("query_scope"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("query_scope"); !isEmptyValue(reflect.ValueOf(queryScopeProp)) && (ok || !reflect.DeepEqual(v, queryScopeProp)) {
		obj["queryScope"] = queryScopeProp
	}
	fieldsProp, err := expandFirestoreIndexFields(d.Get("fields"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("fields"); !isEmptyValue(reflect.ValueOf(fieldsProp)) && (ok || !reflect.DeepEqual(v, fieldsProp)) {
		obj["fields"] = fieldsProp
	}

	return resourceFirestoreIndexEncoder(d, config, obj)
}

func resourceFirestoreIndexEncoder(d TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// We've added project / database / collection as split fields of the name, but
	// the API doesn't expect them.  Make sure we remove them from any requests.

	delete(obj, "project")
	delete(obj, "database")
	delete(obj, "collection")
	return obj, nil
}

func expandFirestoreIndexDatabase(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreIndexCollection(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreIndexQueryScope(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreIndexFields(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedFieldPath, err := expandFirestoreIndexFieldsFieldPath(original["field_path"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedFieldPath); val.IsValid() && !isEmptyValue(val) {
			transformed["fieldPath"] = transformedFieldPath
		}

		transformedOrder, err := expandFirestoreIndexFieldsOrder(original["order"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedOrder); val.IsValid() && !isEmptyValue(val) {
			transformed["order"] = transformedOrder
		}

		transformedArrayConfig, err := expandFirestoreIndexFieldsArrayConfig(original["array_config"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedArrayConfig); val.IsValid() && !isEmptyValue(val) {
			transformed["arrayConfig"] = transformedArrayConfig
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandFirestoreIndexFieldsFieldPath(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreIndexFieldsOrder(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreIndexFieldsArrayConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
