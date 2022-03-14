type User struct {
    gorm.Model
    UUID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
    // ID           uint
    Name         string
    Age          uint8
    Birthday     time.Time
    Foo          string
    Foo2         string
    Email        string
    ActivatedAt  sql.NullTime
    MemberNumber sql.NullString
}
