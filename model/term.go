package model

type Term struct {
	ID        int64  `xorm:"term_id pk"`
	Name      string `xorm:"name"`
	Slug      string `xorm:"slug"`
	TermGroup int64  `xorm:"term_group"`
}

// TableName return the tablename
func (t *Term) TableName() string {
	return "wp_terms"
}

type TermTaxonomy struct {
	ID          int64  `xorm:"term_taxonomy_id pk"`
	TermID      int64  `xorm:"term_id"`
	Taxonomy    string `xorm:"taxonomy"`
	Description string `xorm:"description"`
	Parent      int64  `xorm:"parent"`
	Count       int64  `xorm:"count"`
}

// TableName return the tablename
func (t *TermTaxonomy) TableName() string {
	return "wp_term_taxonomy"
}

type TermRelationship struct {
	ID             int64 `xorm:"object_id"`
	TermTaxonomyID int64 `xorm:"term_taxonomy_id"`
	TermOrder      int   `xorm:"term_order"`
}

// TableName return the tablename
func (t *TermRelationship) TableName() string {
	return "wp_term_relationships"
}

type TermMeta struct {
	ID        int64  `xorm:"meta_id pk"`
	TermID    int64  `xorm:"term_id"`
	MetaKey   string `xorm:"meta_key"`
	MetaValue string `xorm:"meta_value"`
}

// TableName return the tablename
func (ta *TermMeta) TableName() string {
	return "wp_termmeta"
}
