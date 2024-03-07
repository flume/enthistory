// Code generated by ent, DO NOT EDIT.

//go:build tools
// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = "{\"Schema\":\"_examples/basic/ent/schema\",\"Package\":\"_examples/basic/ent\",\"Schemas\":[{\"name\":\"Character\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"friends\",\"type\":\"Character\",\"through\":{\"N\":\"friendships\",\"T\":\"Friendship\"}},{\"name\":\"residence\",\"type\":\"Residence\",\"ref_name\":\"occupants\",\"unique\":true,\"inverse\":true}],\"fields\":[{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"updated_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"age\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"validators\":1,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"name\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"nicknames\",\"type\":{\"Type\":3,\"Ident\":\"[]string\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":true,\"RType\":{\"Name\":\"\",\"Ident\":\"[]string\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":{}}},\"optional\":true,\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"info\",\"type\":{\"Type\":3,\"Ident\":\"map[string]interface {}\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":true,\"RType\":{\"Name\":\"\",\"Ident\":\"map[string]interface {}\",\"Kind\":21,\"PkgPath\":\"\",\"Methods\":{}}},\"optional\":true,\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0}}],\"annotations\":{\"EntSQL\":{\"table\":\"character\"}}},{\"name\":\"CharacterHistory\",\"config\":{\"Table\":\"\"},\"fields\":[{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"updated_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"history_time\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"operation\",\"type\":{\"Type\":6,\"Ident\":\"enthistory.OpType\",\"PkgPath\":\"github.com/flume/enthistory\",\"PkgName\":\"enthistory\",\"Nillable\":false,\"RType\":{\"Name\":\"OpType\",\"Ident\":\"enthistory.OpType\",\"Kind\":24,\"PkgPath\":\"github.com/flume/enthistory\",\"Methods\":{\"MarshalGQL\":{\"In\":[{\"Name\":\"Writer\",\"Ident\":\"io.Writer\",\"Kind\":20,\"PkgPath\":\"io\",\"Methods\":null}],\"Out\":[]},\"String\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"UnmarshalGQL\":{\"In\":[{\"Name\":\"\",\"Ident\":\"interface {}\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Value\":{\"In\":[],\"Out\":[{\"Name\":\"Value\",\"Ident\":\"driver.Value\",\"Kind\":20,\"PkgPath\":\"database/sql/driver\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Values\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]string\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}]}}}},\"enums\":[{\"N\":\"INSERT\",\"V\":\"INSERT\"},{\"N\":\"UPDATE\",\"V\":\"UPDATE\"},{\"N\":\"DELETE\",\"V\":\"DELETE\"}],\"immutable\":true,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"ref\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"immutable\":true,\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"updated_by\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"nillable\":true,\"optional\":true,\"immutable\":true,\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"age\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"immutable\":true,\"position\":{\"Index\":4,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"name\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"immutable\":true,\"position\":{\"Index\":5,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"nicknames\",\"type\":{\"Type\":3,\"Ident\":\"[]string\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":true,\"RType\":{\"Name\":\"\",\"Ident\":\"[]string\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":{}}},\"optional\":true,\"immutable\":true,\"position\":{\"Index\":6,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"info\",\"type\":{\"Type\":3,\"Ident\":\"map[string]interface {}\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":true,\"RType\":{\"Name\":\"\",\"Ident\":\"map[string]interface {}\",\"Kind\":21,\"PkgPath\":\"\",\"Methods\":{}}},\"optional\":true,\"immutable\":true,\"position\":{\"Index\":7,\"MixedIn\":false,\"MixinIndex\":0}}],\"indexes\":[{\"fields\":[\"history_time\"]}],\"annotations\":{\"EntSQL\":{\"table\":\"character_history\"},\"History\":{\"isHistory\":true}}},{\"name\":\"Friendship\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"character\",\"type\":\"Character\",\"field\":\"character_id\",\"unique\":true,\"required\":true},{\"name\":\"friend\",\"type\":\"Character\",\"field\":\"friend_id\",\"unique\":true,\"required\":true}],\"fields\":[{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"updated_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"id\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"character_id\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"friend_id\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0}}],\"annotations\":{\"EntSQL\":{\"table\":\"friendship\"}}},{\"name\":\"FriendshipHistory\",\"config\":{\"Table\":\"\"},\"fields\":[{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"updated_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"id\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"history_time\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"operation\",\"type\":{\"Type\":6,\"Ident\":\"enthistory.OpType\",\"PkgPath\":\"github.com/flume/enthistory\",\"PkgName\":\"enthistory\",\"Nillable\":false,\"RType\":{\"Name\":\"OpType\",\"Ident\":\"enthistory.OpType\",\"Kind\":24,\"PkgPath\":\"github.com/flume/enthistory\",\"Methods\":{\"MarshalGQL\":{\"In\":[{\"Name\":\"Writer\",\"Ident\":\"io.Writer\",\"Kind\":20,\"PkgPath\":\"io\",\"Methods\":null}],\"Out\":[]},\"String\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"UnmarshalGQL\":{\"In\":[{\"Name\":\"\",\"Ident\":\"interface {}\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Value\":{\"In\":[],\"Out\":[{\"Name\":\"Value\",\"Ident\":\"driver.Value\",\"Kind\":20,\"PkgPath\":\"database/sql/driver\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Values\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]string\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}]}}}},\"enums\":[{\"N\":\"INSERT\",\"V\":\"INSERT\"},{\"N\":\"UPDATE\",\"V\":\"UPDATE\"},{\"N\":\"DELETE\",\"V\":\"DELETE\"}],\"immutable\":true,\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"ref\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"immutable\":true,\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"updated_by\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"nillable\":true,\"optional\":true,\"immutable\":true,\"position\":{\"Index\":4,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"character_id\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"immutable\":true,\"position\":{\"Index\":5,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"friend_id\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"immutable\":true,\"position\":{\"Index\":6,\"MixedIn\":false,\"MixinIndex\":0}}],\"indexes\":[{\"fields\":[\"history_time\"]}],\"annotations\":{\"EntSQL\":{\"table\":\"friendship_history\"},\"History\":{\"isHistory\":true}}},{\"name\":\"Residence\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"occupants\",\"type\":\"Character\"}],\"fields\":[{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"updated_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"id\",\"type\":{\"Type\":4,\"Ident\":\"uuid.UUID\",\"PkgPath\":\"github.com/google/uuid\",\"PkgName\":\"uuid\",\"Nillable\":false,\"RType\":{\"Name\":\"UUID\",\"Ident\":\"uuid.UUID\",\"Kind\":17,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":{\"ClockSequence\":{\"In\":[],\"Out\":[{\"Name\":\"int\",\"Ident\":\"int\",\"Kind\":2,\"PkgPath\":\"\",\"Methods\":null}]},\"Domain\":{\"In\":[],\"Out\":[{\"Name\":\"Domain\",\"Ident\":\"uuid.Domain\",\"Kind\":8,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]},\"ID\":{\"In\":[],\"Out\":[{\"Name\":\"uint32\",\"Ident\":\"uint32\",\"Kind\":10,\"PkgPath\":\"\",\"Methods\":null}]},\"MarshalBinary\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"MarshalText\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"NodeID\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}]},\"Scan\":{\"In\":[{\"Name\":\"\",\"Ident\":\"interface {}\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"String\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"Time\":{\"In\":[],\"Out\":[{\"Name\":\"Time\",\"Ident\":\"uuid.Time\",\"Kind\":6,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]},\"URN\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"UnmarshalBinary\":{\"In\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"UnmarshalText\":{\"In\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Value\":{\"In\":[],\"Out\":[{\"Name\":\"Value\",\"Ident\":\"driver.Value\",\"Kind\":20,\"PkgPath\":\"database/sql/driver\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Variant\":{\"In\":[],\"Out\":[{\"Name\":\"Variant\",\"Ident\":\"uuid.Variant\",\"Kind\":8,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]},\"Version\":{\"In\":[],\"Out\":[{\"Name\":\"Version\",\"Ident\":\"uuid.Version\",\"Kind\":8,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]}}}},\"default\":true,\"default_kind\":19,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"name\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}}],\"annotations\":{\"EntSQL\":{\"table\":\"residence\"}}},{\"name\":\"ResidenceHistory\",\"config\":{\"Table\":\"\"},\"fields\":[{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"updated_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"id\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"history_time\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"operation\",\"type\":{\"Type\":6,\"Ident\":\"enthistory.OpType\",\"PkgPath\":\"github.com/flume/enthistory\",\"PkgName\":\"enthistory\",\"Nillable\":false,\"RType\":{\"Name\":\"OpType\",\"Ident\":\"enthistory.OpType\",\"Kind\":24,\"PkgPath\":\"github.com/flume/enthistory\",\"Methods\":{\"MarshalGQL\":{\"In\":[{\"Name\":\"Writer\",\"Ident\":\"io.Writer\",\"Kind\":20,\"PkgPath\":\"io\",\"Methods\":null}],\"Out\":[]},\"String\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"UnmarshalGQL\":{\"In\":[{\"Name\":\"\",\"Ident\":\"interface {}\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Value\":{\"In\":[],\"Out\":[{\"Name\":\"Value\",\"Ident\":\"driver.Value\",\"Kind\":20,\"PkgPath\":\"database/sql/driver\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Values\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]string\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}]}}}},\"enums\":[{\"N\":\"INSERT\",\"V\":\"INSERT\"},{\"N\":\"UPDATE\",\"V\":\"UPDATE\"},{\"N\":\"DELETE\",\"V\":\"DELETE\"}],\"immutable\":true,\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"ref\",\"type\":{\"Type\":4,\"Ident\":\"uuid.UUID\",\"PkgPath\":\"github.com/google/uuid\",\"PkgName\":\"uuid\",\"Nillable\":false,\"RType\":{\"Name\":\"UUID\",\"Ident\":\"uuid.UUID\",\"Kind\":17,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":{\"ClockSequence\":{\"In\":[],\"Out\":[{\"Name\":\"int\",\"Ident\":\"int\",\"Kind\":2,\"PkgPath\":\"\",\"Methods\":null}]},\"Domain\":{\"In\":[],\"Out\":[{\"Name\":\"Domain\",\"Ident\":\"uuid.Domain\",\"Kind\":8,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]},\"ID\":{\"In\":[],\"Out\":[{\"Name\":\"uint32\",\"Ident\":\"uint32\",\"Kind\":10,\"PkgPath\":\"\",\"Methods\":null}]},\"MarshalBinary\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"MarshalText\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"NodeID\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}]},\"Scan\":{\"In\":[{\"Name\":\"\",\"Ident\":\"interface {}\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"String\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"Time\":{\"In\":[],\"Out\":[{\"Name\":\"Time\",\"Ident\":\"uuid.Time\",\"Kind\":6,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]},\"URN\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"UnmarshalBinary\":{\"In\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"UnmarshalText\":{\"In\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Value\":{\"In\":[],\"Out\":[{\"Name\":\"Value\",\"Ident\":\"driver.Value\",\"Kind\":20,\"PkgPath\":\"database/sql/driver\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Variant\":{\"In\":[],\"Out\":[{\"Name\":\"Variant\",\"Ident\":\"uuid.Variant\",\"Kind\":8,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]},\"Version\":{\"In\":[],\"Out\":[{\"Name\":\"Version\",\"Ident\":\"uuid.Version\",\"Kind\":8,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]}}}},\"optional\":true,\"immutable\":true,\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"updated_by\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"nillable\":true,\"optional\":true,\"immutable\":true,\"position\":{\"Index\":4,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"name\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"immutable\":true,\"position\":{\"Index\":5,\"MixedIn\":false,\"MixinIndex\":0}}],\"indexes\":[{\"fields\":[\"history_time\"]}],\"annotations\":{\"EntSQL\":{\"table\":\"residence_history\"},\"History\":{\"isHistory\":true}}}],\"Features\":[\"schema/snapshot\"]}"
