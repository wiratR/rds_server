package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CardType struct for storing card details with GORM annotations
type CardType struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"` // UUID as primary key with default generation
	CardId      int       `gorm:"type:int;uniqueIndex;not null"`                    // Card ID, unique and not null
	ShortName   string    `gorm:"type:varchar(5);uniqueIndex;not null"`             // Card Short Name, unique and not null
	Description string    `gorm:"type:varchar(100);uniqueIndex;not null"`           // Description, unique and not null
	CreatedAt   time.Time // Automatically handled by GORM for creation timestamp
	UpdatedAt   time.Time // Automatically handled by GORM for update timestamp
}

// MediaType struct for storing media type details with GORM annotations
type MediaType struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"` // UUID as primary key with default generation
	MediaTypeId int       `gorm:"type:int;uniqueIndex;not null"`                    // Media Type ID, unique and not null
	ShortName   string    `gorm:"type:varchar(5);uniqueIndex;not null"`             // Media Type Short Name, unique and not null
	Description string    `gorm:"type:varchar(100);uniqueIndex;not null"`           // Description, unique and not null
	CreatedAt   time.Time // Automatically handled by GORM for creation timestamp
	UpdatedAt   time.Time // Automatically handled by GORM for update timestamp
}

// Station struct for storing station details with GORM annotations
type Station struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"` // UUID as primary key with default generation
	StationId   int       `gorm:"type:int;uniqueIndex;not null"`                    // Station ID, unique and not null
	ShortName   string    `gorm:"type:varchar(5);uniqueIndex;not null"`             // Station Short Name, unique and not null
	Description string    `gorm:"type:varchar(100);uniqueIndex;not null"`           // Description, unique and not null
	IsCrossLine bool      `gorm:"type:bool;default:false;not null"`                 // IsCrossLine, default is false
	CreatedAt   time.Time // Automatically handled by GORM for creation timestamp
	UpdatedAt   time.Time // Automatically handled by GORM for update timestamp
}

// Line struct for storing line details with GORM annotations
type Line struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"` // UUID as primary key with default generation
	LineId      int       `gorm:"type:int;uniqueIndex;not null"`                    // Line ID, unique and not null
	ShortName   string    `gorm:"type:varchar(5);uniqueIndex;not null"`             // Line Short Name, unique and not null
	Description string    `gorm:"type:varchar(100);uniqueIndex;not null"`           // Description, unique and not null
	CreatedAt   time.Time // Automatically handled by GORM for creation timestamp
	UpdatedAt   time.Time // Automatically handled by GORM for update timestamp
}

// ServiceProvider struct for storing service provider details with GORM annotations
type ServiceProvider struct {
	ID                uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	ServiceProviderId int       `gorm:"type:int;uniqueIndex;not null"`          // Service Provider ID, unique and not null
	ShortName         string    `gorm:"type:varchar(5);uniqueIndex;not null"`   // Service Provider Short Name, unique and not null
	Description       string    `gorm:"type:varchar(100);uniqueIndex;not null"` // Description, unique and not null
	CreatedAt         time.Time // Automatically handled by GORM for creation timestamp
	UpdatedAt         time.Time // Automatically handled by GORM for update timestamp
}

// Fare struct for storing fare details with GORM annotations
type Fare struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	SpId        int       `gorm:"type:int;not null"`
	LineId      int       `gorm:"type:int;not null"`
	StationId   int       `gorm:"type:int;not null"`
	CardTypeId  int       `gorm:"type:int;not null"`
	MediaTypeId int       `gorm:"type:int;not null"`
	Amount      int       `gorm:"type:int;not null"`
	CreatedAt   time.Time // Automatically handled by GORM for creation timestamp
	UpdatedAt   time.Time // Automatically handled by GORM for update timestamp
}

// Ensure to initialize the UUID fields before creating records, if necessary
func (c *CardType) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return
}

func (m *MediaType) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return
}

func (s *Station) BeforeCreate(tx *gorm.DB) (err error) {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
	return
}

func (l *Line) BeforeCreate(tx *gorm.DB) (err error) {
	if l.ID == uuid.Nil {
		l.ID = uuid.New()
	}
	return
}

func (sp *ServiceProvider) BeforeCreate(tx *gorm.DB) (err error) {
	if sp.ID == uuid.Nil {
		sp.ID = uuid.New()
	}
	return
}

func (f *Fare) BeforeCreate(tx *gorm.DB) (err error) {
	if f.ID == uuid.Nil {
		f.ID = uuid.New()
	}
	return
}
