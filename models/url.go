type URL struct {
    gorm.Model
    //URL_id       int64  `gorm:"primaryKey"`
    Disabled     bool
    UUID         string `gorm:"type:uuid;default:uuid_generate_v4()"`
    Note         string
    URL_group_id int
    Path          string `gorm:"unique;not null"`
    Rsp_code     int
    Rsp_code_exp int `gorm:"default:200`
    Rsp_code_test   bool
    Rsp_time     float64 `gorm:"type:decimal(16,6);default:0"`
    Rsp_time_exp int `gorm:"default:4`
    Rsp_time_test   bool
    Rsp_regex_exp string `gorm:"default:statushealthy`
    Rsp_regex_test   bool
    AllowInsecureTLS    bool `gorm:"default:false`
    // Rsp_time     Decimal `gorm:"type:decimal(16,6);default:0"`
    //Amount       float32   `sql:"type:decimal(10,2);"`
    Sequence     int
    // Rsp_time     float64
}
