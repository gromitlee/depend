package rel

// Relation 资源A 被 资源B 依赖
type Relation struct {
	// 资源A id
	IDA string `gorm:"primaryKey;index:idx_relation_id_a;autoIncrement:false;not null"`
	// 资源A 类型
	TypA ResType `gorm:"primaryKey;index:idx_relation_typ_a;autoIncrement:false;not null"`
	// 依赖方式
	RelTyp Type `gorm:"primaryKey;index:idx_relation_rel_typ;autoIncrement:false;not null"`
	// 资源B id
	IDB string `gorm:"primaryKey;index:idx_relation_id_b;autoIncrement:false;not null"`
	// 资源A 类型
	TypB ResType `gorm:"primaryKey;index:idx_relation_typ_b;autoIncrement:false;not null"`
}
