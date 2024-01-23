package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Metric holds the schema definition for the Metric entity.
type Metric struct {
	ent.Schema
}

// Fields of the Metric.
func (Metric) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("uuid"),
		field.Enum("metric_type").
			Values(
				"calories",       // Калории
				"fats",           // Жиры
				"carbs",          // Углеводы
				"proteins",       // Белки
				"vegetable",      // Растительная пища
				"steps",          // Шаги
				"weight",         // Вес
				"neck",           // Шея
				"breast",         // Грудь
				"left_biceps",    // Левый бицепс
				"right_biceps",   // Правый бицепс
				"waist",          // Талия
				"buttocks",       // Ягодицы
				"left_hip",       // Левое бедро
				"right_hip",      // Правое бедро
				"left_calf",      // Левая икра
				"right_calf",     // Правая икра
				"fat_percentage", // Процент жира
			),
		field.Int("value"),
		field.Time("date"),
	}
}

// Edges of the Metric.
func (Metric) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("metrics").
			Unique(),
	}
}

// Indexes of the Metric.
func (Metric) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("metric_type").
			Edges("owner"),
	}
}
