package admin.authz

# user-role assignments
user_roles := {
	"business": ["permission_editer", "product_editer", "order_editer", "customer_editer", "staff_editer", "business_editer", "sales_editer", "report_viewer"],
	"business_agent": ["product_editer", "order_editer", "customer_editer", "staff_editer", "sales_editer", "report_viewer"],
	"staff": ["permission_editer", "product_editer", "order_editer", "customer_editer", "staff_editer", "sales_editer", "report_viewer"],
	"staff_agent": ["product_editer", "order_editer", "customer_editer", "staff_editer", "sales_editer", "report_viewer"],
	"sales": ["order_viewer", "customer_viewer", "report_viewer", "product_viewer"],
	"customer": ["product_viewer"],
}

# role-permission assignments
role_permissions := {
	"self": [
		{"domain": "self", "package": "self", "action": "view"},
		{"domain": "self", "package": "self", "action": "edit"},
	],
	"admin": [
		{"domain": "resource", "package": "resource", "action": "view"},
		{"domain": "resource", "package": "resource", "action": "edit"},
	],
	"permission_editer": [
		{"domain": "permission", "package": "permission", "action": "view"},
		{"domain": "permission", "package": "permission", "action": "edit"},
	],
	"product_editer": [
		{"domain": "product", "package": "product", "action": "view"},
		{"domain": "product", "package": "product", "action": "edit"},
	],
	"report_viewer": [{"domain": "report", "package": "report", "action": "view"}],
	"product_viewer": [{"domain": "product", "package": "product", "action": "view"}],
	"order_editer": [
		{"domain": "order", "package": "order", "action": "view"},
		{"domain": "order", "package": "order", "action": "edit"},
	],
	"order_viewer": [{"domain": "order", "package": "order", "action": "view"}],
	"customer_editer": [
		{"domain": "customer", "package": "customer", "action": "view"},
		{"domain": "customer", "package": "customer", "action": "edit"},
	],
	"customer_viewer": [{"domain": "customer", "package": "customer", "action": "view"}],
	"staff_editer": [
		{"domain": "staff", "package": "staff", "action": "view"},
		{"domain": "staff", "package": "staff", "action": "edit"},
	],
	"staff_viewer": [{"domain": "staff", "package": "staff", "action": "view"}],
	"business_editer": [
		{"domain": "business", "package": "business", "action": "view"},
		{"domain": "business", "package": "business", "action": "edit"},
	],
	"business_viewer": [{"domain": "business", "package": "business", "action": "view"}],
	"sales_editer": [
		{"domain": "sales", "package": "sales", "action": "view"},
		{"domain": "sales", "package": "sales", "action": "edit"},
	],
	"sales_viewer": [{"domain": "sales", "package": "sales", "action": "view"}],
}
