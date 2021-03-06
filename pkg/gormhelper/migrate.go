// Copyright © 2020 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gormhelper

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func AddForeignKeyAndReferencedKey(db *gorm.DB, logger logrus.FieldLogger, parentTable, childTable interface{}, foreignKeyField string, referencedField string) error {
	parentTableScope := db.NewScope(parentTable)
	childTableScope := db.NewScope(childTable)

	log := logger.WithFields(logrus.Fields{
		"parent_table": strings.TrimSpace(parentTableScope.TableName()),
		"child_table":  strings.TrimSpace(childTableScope.TableName()),
	})

	f, ok := childTableScope.FieldByName(foreignKeyField)
	if !ok {
		return fmt.Errorf("field %q not found", foreignKeyField)
	}
	if !f.IsForeignKey {
		return fmt.Errorf("%q is not a foreign key field", foreignKeyField)
	}

	parentIdField := ""
	if referencedField == "" {
		parentIdField = parentTableScope.PrimaryKey()
	} else {
		f, ok := parentTableScope.FieldByName(referencedField)
		if !ok {
			return fmt.Errorf("field %q not found", referencedField)
		}
		parentIdField = f.DBName
	}
	references := fmt.Sprintf("%s(%s)", parentTableScope.TableName(), parentIdField)

	log.Infof("adding foreign key constraint: %s -> %s", f.DBName, references)
	return db.Model(childTable).AddForeignKey(f.DBName, references, "RESTRICT", "RESTRICT").Error
}

func AddForeignKey(db *gorm.DB, logger logrus.FieldLogger, parentTable, childTable interface{}, foreignKeyField string) error {
	return AddForeignKeyAndReferencedKey(db, logger, parentTable, childTable, foreignKeyField, "")
}
