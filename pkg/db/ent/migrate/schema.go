// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AppGoodsColumns holds the columns for the "app_goods" table.
	AppGoodsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "good_id", Type: field.TypeUUID},
		{Name: "online", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "visible", Type: field.TypeBool, Nullable: true, Default: true},
		{Name: "good_name", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "price", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "display_index", Type: field.TypeInt32, Nullable: true, Default: 0},
		{Name: "purchase_limit", Type: field.TypeInt32, Nullable: true, Default: 3000},
		{Name: "commission_percent", Type: field.TypeInt32, Nullable: true, Default: 0},
		{Name: "sale_start_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "sale_end_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "service_start_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "technical_fee_ratio", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "electricity_fee_ratio", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "daily_reward_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "commission_settle_type", Type: field.TypeString, Nullable: true, Default: "NoCommission"},
	}
	// AppGoodsTable holds the schema information for the "app_goods" table.
	AppGoodsTable = &schema.Table{
		Name:       "app_goods",
		Columns:    AppGoodsColumns,
		PrimaryKey: []*schema.Column{AppGoodsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "appgood_good_id_app_id_online",
				Unique:  false,
				Columns: []*schema.Column{AppGoodsColumns[5], AppGoodsColumns[4], AppGoodsColumns[6]},
			},
		},
	}
	// CommentsColumns holds the columns for the "comments" table.
	CommentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "good_id", Type: field.TypeUUID},
		{Name: "order_id", Type: field.TypeUUID},
		{Name: "content", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "reply_to_id", Type: field.TypeUUID, Nullable: true},
	}
	// CommentsTable holds the schema information for the "comments" table.
	CommentsTable = &schema.Table{
		Name:       "comments",
		Columns:    CommentsColumns,
		PrimaryKey: []*schema.Column{CommentsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "comment_app_id_user_id_good_id_order_id",
				Unique:  false,
				Columns: []*schema.Column{CommentsColumns[4], CommentsColumns[5], CommentsColumns[6], CommentsColumns[7]},
			},
		},
	}
	// DeviceInfosColumns holds the columns for the "device_infos" table.
	DeviceInfosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "type", Type: field.TypeString, Nullable: true, Size: 64, Default: ""},
		{Name: "manufacturer", Type: field.TypeString, Nullable: true, Size: 64, Default: ""},
		{Name: "power_comsuption", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "shipment_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "posters", Type: field.TypeJSON, Nullable: true},
	}
	// DeviceInfosTable holds the schema information for the "device_infos" table.
	DeviceInfosTable = &schema.Table{
		Name:       "device_infos",
		Columns:    DeviceInfosColumns,
		PrimaryKey: []*schema.Column{DeviceInfosColumns[0]},
	}
	// ExtraInfosColumns holds the columns for the "extra_infos" table.
	ExtraInfosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "good_id", Type: field.TypeUUID},
		{Name: "posters", Type: field.TypeJSON, Nullable: true},
		{Name: "labels", Type: field.TypeJSON, Nullable: true},
		{Name: "vote_count", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "rating", Type: field.TypeFloat32, Nullable: true, Default: 0},
	}
	// ExtraInfosTable holds the schema information for the "extra_infos" table.
	ExtraInfosTable = &schema.Table{
		Name:       "extra_infos",
		Columns:    ExtraInfosColumns,
		PrimaryKey: []*schema.Column{ExtraInfosColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "extrainfo_good_id",
				Unique:  false,
				Columns: []*schema.Column{ExtraInfosColumns[4]},
			},
		},
	}
	// GoodsColumns holds the columns for the "goods" table.
	GoodsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "device_info_id", Type: field.TypeUUID},
		{Name: "duration_days", Type: field.TypeInt32, Nullable: true, Default: 365},
		{Name: "coin_type_id", Type: field.TypeUUID},
		{Name: "inherit_from_good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "vendor_location_id", Type: field.TypeUUID},
		{Name: "price", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "benefit_type", Type: field.TypeString, Nullable: true, Default: "DefaultBenefitType"},
		{Name: "good_type", Type: field.TypeString, Nullable: true, Default: "DefaultGoodType"},
		{Name: "title", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "unit", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "unit_amount", Type: field.TypeInt32, Nullable: true, Default: 0},
		{Name: "support_coin_type_ids", Type: field.TypeJSON, Nullable: true},
		{Name: "delivery_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "start_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "test_only", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "benefit_interval_hours", Type: field.TypeUint32, Nullable: true, Default: 24},
		{Name: "benefit_state", Type: field.TypeString, Nullable: true, Default: "BenefitWait"},
		{Name: "last_benefit_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "benefit_tids", Type: field.TypeJSON, Nullable: true},
		{Name: "next_benefit_start_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
	}
	// GoodsTable holds the schema information for the "goods" table.
	GoodsTable = &schema.Table{
		Name:       "goods",
		Columns:    GoodsColumns,
		PrimaryKey: []*schema.Column{GoodsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "good_device_info_id_vendor_location_id_benefit_type_good_type",
				Unique:  false,
				Columns: []*schema.Column{GoodsColumns[4], GoodsColumns[8], GoodsColumns[10], GoodsColumns[11]},
			},
		},
	}
	// PromotionsColumns holds the columns for the "promotions" table.
	PromotionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "good_id", Type: field.TypeUUID},
		{Name: "message", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "start_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "end_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "price", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "posters", Type: field.TypeJSON, Nullable: true},
	}
	// PromotionsTable holds the schema information for the "promotions" table.
	PromotionsTable = &schema.Table{
		Name:       "promotions",
		Columns:    PromotionsColumns,
		PrimaryKey: []*schema.Column{PromotionsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "promotion_good_id_app_id",
				Unique:  false,
				Columns: []*schema.Column{PromotionsColumns[5], PromotionsColumns[4]},
			},
		},
	}
	// RecommendsColumns holds the columns for the "recommends" table.
	RecommendsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "good_id", Type: field.TypeUUID},
		{Name: "recommender_id", Type: field.TypeUUID, Nullable: true},
		{Name: "message", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "recommend_index", Type: field.TypeFloat64, Nullable: true, Default: 0},
	}
	// RecommendsTable holds the schema information for the "recommends" table.
	RecommendsTable = &schema.Table{
		Name:       "recommends",
		Columns:    RecommendsColumns,
		PrimaryKey: []*schema.Column{RecommendsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "recommend_app_id_good_id",
				Unique:  false,
				Columns: []*schema.Column{RecommendsColumns[4], RecommendsColumns[5]},
			},
		},
	}
	// StocksColumns holds the columns for the "stocks" table.
	StocksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "good_id", Type: field.TypeUUID},
		{Name: "total", Type: field.TypeUint32},
		{Name: "locked", Type: field.TypeUint32},
		{Name: "in_service", Type: field.TypeUint32},
		{Name: "wait_start", Type: field.TypeUint32},
		{Name: "sold", Type: field.TypeUint32},
	}
	// StocksTable holds the schema information for the "stocks" table.
	StocksTable = &schema.Table{
		Name:       "stocks",
		Columns:    StocksColumns,
		PrimaryKey: []*schema.Column{StocksColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "stock_good_id",
				Unique:  true,
				Columns: []*schema.Column{StocksColumns[4]},
			},
		},
	}
	// SubGoodsColumns holds the columns for the "sub_goods" table.
	SubGoodsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "main_good_id", Type: field.TypeUUID},
		{Name: "sub_good_id", Type: field.TypeUUID},
		{Name: "must", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "commission", Type: field.TypeBool, Nullable: true, Default: false},
	}
	// SubGoodsTable holds the schema information for the "sub_goods" table.
	SubGoodsTable = &schema.Table{
		Name:       "sub_goods",
		Columns:    SubGoodsColumns,
		PrimaryKey: []*schema.Column{SubGoodsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "subgood_app_id_main_good_id_sub_good_id",
				Unique:  false,
				Columns: []*schema.Column{SubGoodsColumns[4], SubGoodsColumns[5], SubGoodsColumns[6]},
			},
		},
	}
	// VendorLocationsColumns holds the columns for the "vendor_locations" table.
	VendorLocationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "country", Type: field.TypeString, Nullable: true, Size: 128, Default: ""},
		{Name: "province", Type: field.TypeString, Nullable: true, Size: 128, Default: ""},
		{Name: "city", Type: field.TypeString, Nullable: true, Size: 128, Default: ""},
		{Name: "address", Type: field.TypeString, Nullable: true, Size: 256, Default: ""},
	}
	// VendorLocationsTable holds the schema information for the "vendor_locations" table.
	VendorLocationsTable = &schema.Table{
		Name:       "vendor_locations",
		Columns:    VendorLocationsColumns,
		PrimaryKey: []*schema.Column{VendorLocationsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AppGoodsTable,
		CommentsTable,
		DeviceInfosTable,
		ExtraInfosTable,
		GoodsTable,
		PromotionsTable,
		RecommendsTable,
		StocksTable,
		SubGoodsTable,
		VendorLocationsTable,
	}
)

func init() {
}
