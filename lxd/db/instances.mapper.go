//go:build linux && cgo && !agent
// +build linux,cgo,!agent

package db

// The code below was generated by lxd-generate - DO NOT EDIT!

import (
	"database/sql"
	"fmt"

	"github.com/lxc/lxd/lxd/db/cluster"
	"github.com/lxc/lxd/lxd/db/query"
	"github.com/lxc/lxd/shared/api"
	"github.com/pkg/errors"
)

var _ = api.ServerEnvironment{}

var instanceNames = cluster.RegisterStmt(`
SELECT projects.name AS project, instances.name
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  ORDER BY projects.id, instances.name
`)

var instanceNamesByProject = cluster.RegisterStmt(`
SELECT projects.name AS project, instances.name
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE project = ? ORDER BY projects.id, instances.name
`)

var instanceNamesByID = cluster.RegisterStmt(`
SELECT projects.name AS project, instances.name
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE instances.id = ? ORDER BY projects.id, instances.name
`)

var instanceObjects = cluster.RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  ORDER BY projects.id, instances.name
`)

var instanceObjectsByProject = cluster.RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE project = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByProjectAndType = cluster.RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE project = ? AND instances.type = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByProjectAndTypeAndNode = cluster.RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE project = ? AND instances.type = ? AND node = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByProjectAndTypeAndNodeAndName = cluster.RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE project = ? AND instances.type = ? AND node = ? AND instances.name = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByProjectAndTypeAndName = cluster.RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE project = ? AND instances.type = ? AND instances.name = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByProjectAndName = cluster.RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE project = ? AND instances.name = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByProjectAndNameAndNode = cluster.RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE project = ? AND instances.name = ? AND node = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByProjectAndNode = cluster.RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE project = ? AND node = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByType = cluster.RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE instances.type = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByTypeAndName = cluster.RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE instances.type = ? AND instances.name = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByTypeAndNameAndNode = cluster.RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE instances.type = ? AND instances.name = ? AND node = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByTypeAndNode = cluster.RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE instances.type = ? AND node = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByNode = cluster.RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE node = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByNodeAndName = cluster.RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE node = ? AND instances.name = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByName = cluster.RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE instances.name = ? ORDER BY projects.id, instances.name
`)

var instanceID = cluster.RegisterStmt(`
SELECT instances.id FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE projects.name = ? AND instances.name = ?
`)

var instanceCreate = cluster.RegisterStmt(`
INSERT INTO instances (project_id, name, node_id, type, architecture, ephemeral, creation_date, stateful, last_use_date, description, expiry_date)
  VALUES ((SELECT projects.id FROM projects WHERE projects.name = ?), ?, (SELECT nodes.id FROM nodes WHERE nodes.name = ?), ?, ?, ?, ?, ?, ?, ?, ?)
`)

var instanceRename = cluster.RegisterStmt(`
UPDATE instances SET name = ? WHERE project_id = (SELECT projects.id FROM projects WHERE projects.name = ?) AND name = ?
`)

var instanceDeleteByProjectAndName = cluster.RegisterStmt(`
DELETE FROM instances WHERE project_id = (SELECT projects.id FROM projects WHERE projects.name = ?) AND name = ?
`)

var instanceUpdate = cluster.RegisterStmt(`
UPDATE instances
  SET project_id = (SELECT id FROM projects WHERE name = ?), name = ?, node_id = (SELECT id FROM nodes WHERE name = ?), type = ?, architecture = ?, ephemeral = ?, creation_date = ?, stateful = ?, last_use_date = ?, description = ?, expiry_date = ?
 WHERE id = ?
`)

// GetInstances returns all available instances.
// generator: instance GetMany
func (c *ClusterTx) GetInstances(filter InstanceFilter) ([]Instance, error) {
	// Result slice.
	objects := make([]Instance, 0)

	// Pick the prepared statement and arguments to use based on active criteria.
	var stmt *sql.Stmt
	var args []interface{}

	if filter.Project != nil && filter.Type != nil && filter.Node != nil && filter.Name != nil && filter.ID == nil {
		stmt = c.stmt(instanceObjectsByProjectAndTypeAndNodeAndName)
		args = []interface{}{
			filter.Project,
			filter.Type,
			filter.Node,
			filter.Name,
		}
	} else if filter.Project != nil && filter.Type != nil && filter.Node != nil && filter.ID == nil && filter.Name == nil {
		stmt = c.stmt(instanceObjectsByProjectAndTypeAndNode)
		args = []interface{}{
			filter.Project,
			filter.Type,
			filter.Node,
		}
	} else if filter.Project != nil && filter.Type != nil && filter.Name != nil && filter.ID == nil && filter.Node == nil {
		stmt = c.stmt(instanceObjectsByProjectAndTypeAndName)
		args = []interface{}{
			filter.Project,
			filter.Type,
			filter.Name,
		}
	} else if filter.Type != nil && filter.Name != nil && filter.Node != nil && filter.ID == nil && filter.Project == nil {
		stmt = c.stmt(instanceObjectsByTypeAndNameAndNode)
		args = []interface{}{
			filter.Type,
			filter.Name,
			filter.Node,
		}
	} else if filter.Project != nil && filter.Name != nil && filter.Node != nil && filter.ID == nil && filter.Type == nil {
		stmt = c.stmt(instanceObjectsByProjectAndNameAndNode)
		args = []interface{}{
			filter.Project,
			filter.Name,
			filter.Node,
		}
	} else if filter.Project != nil && filter.Type != nil && filter.ID == nil && filter.Name == nil && filter.Node == nil {
		stmt = c.stmt(instanceObjectsByProjectAndType)
		args = []interface{}{
			filter.Project,
			filter.Type,
		}
	} else if filter.Type != nil && filter.Node != nil && filter.ID == nil && filter.Project == nil && filter.Name == nil {
		stmt = c.stmt(instanceObjectsByTypeAndNode)
		args = []interface{}{
			filter.Type,
			filter.Node,
		}
	} else if filter.Type != nil && filter.Name != nil && filter.ID == nil && filter.Project == nil && filter.Node == nil {
		stmt = c.stmt(instanceObjectsByTypeAndName)
		args = []interface{}{
			filter.Type,
			filter.Name,
		}
	} else if filter.Project != nil && filter.Node != nil && filter.ID == nil && filter.Name == nil && filter.Type == nil {
		stmt = c.stmt(instanceObjectsByProjectAndNode)
		args = []interface{}{
			filter.Project,
			filter.Node,
		}
	} else if filter.Project != nil && filter.Name != nil && filter.ID == nil && filter.Node == nil && filter.Type == nil {
		stmt = c.stmt(instanceObjectsByProjectAndName)
		args = []interface{}{
			filter.Project,
			filter.Name,
		}
	} else if filter.Node != nil && filter.Name != nil && filter.ID == nil && filter.Project == nil && filter.Type == nil {
		stmt = c.stmt(instanceObjectsByNodeAndName)
		args = []interface{}{
			filter.Node,
			filter.Name,
		}
	} else if filter.Type != nil && filter.ID == nil && filter.Project == nil && filter.Name == nil && filter.Node == nil {
		stmt = c.stmt(instanceObjectsByType)
		args = []interface{}{
			filter.Type,
		}
	} else if filter.Project != nil && filter.ID == nil && filter.Name == nil && filter.Node == nil && filter.Type == nil {
		stmt = c.stmt(instanceObjectsByProject)
		args = []interface{}{
			filter.Project,
		}
	} else if filter.Node != nil && filter.ID == nil && filter.Project == nil && filter.Name == nil && filter.Type == nil {
		stmt = c.stmt(instanceObjectsByNode)
		args = []interface{}{
			filter.Node,
		}
	} else if filter.Name != nil && filter.ID == nil && filter.Project == nil && filter.Node == nil && filter.Type == nil {
		stmt = c.stmt(instanceObjectsByName)
		args = []interface{}{
			filter.Name,
		}
	} else if filter.ID == nil && filter.Project == nil && filter.Name == nil && filter.Node == nil && filter.Type == nil {
		stmt = c.stmt(instanceObjects)
		args = []interface{}{}
	} else {
		return nil, fmt.Errorf("No statement exists for the given Filter")
	}

	// Dest function for scanning a row.
	dest := func(i int) []interface{} {
		objects = append(objects, Instance{})
		return []interface{}{
			&objects[i].ID,
			&objects[i].Project,
			&objects[i].Name,
			&objects[i].Node,
			&objects[i].Type,
			&objects[i].Architecture,
			&objects[i].Ephemeral,
			&objects[i].CreationDate,
			&objects[i].Stateful,
			&objects[i].LastUseDate,
			&objects[i].Description,
			&objects[i].ExpiryDate,
		}
	}

	// Select.
	err := query.SelectObjects(stmt, dest, args...)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to fetch instances")
	}

	return objects, nil
}

// GetInstance returns the instance with the given key.
// generator: instance GetOne
func (c *ClusterTx) GetInstance(project string, name string) (*Instance, error) {
	filter := InstanceFilter{}
	filter.Project = &project
	filter.Name = &name

	objects, err := c.GetInstances(filter)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to fetch Instance")
	}

	switch len(objects) {
	case 0:
		return nil, ErrNoSuchObject
	case 1:
		return &objects[0], nil
	default:
		return nil, fmt.Errorf("More than one instance matches")
	}
}

// GetInstanceURIs returns all available instance URIs.
// generator: instance URIs
func (c *ClusterTx) GetInstanceURIs(filter InstanceFilter) ([]string, error) {
	var args []interface{}
	var stmt *sql.Stmt
	if filter.Project != nil && filter.ID == nil && filter.Name == nil && filter.Type == nil {
		stmt = c.stmt(instanceNamesByProject)
		args = []interface{}{
			filter.Project,
		}
	} else if filter.ID != nil && filter.Project == nil && filter.Name == nil && filter.Type == nil {
		stmt = c.stmt(instanceNamesByID)
		args = []interface{}{
			filter.ID,
		}
	} else if filter.ID == nil && filter.Project == nil && filter.Name == nil && filter.Type == nil {
		stmt = c.stmt(instanceNames)
		args = []interface{}{}
	} else {
		return nil, fmt.Errorf("No statement exists for the given Filter")
	}

	code := cluster.EntityTypes["instance"]
	formatter := cluster.EntityFormatURIs[code]

	return query.SelectURIs(stmt, formatter, args...)
}

// GetInstanceID return the ID of the instance with the given key.
// generator: instance ID
func (c *ClusterTx) GetInstanceID(project string, name string) (int64, error) {
	stmt := c.stmt(instanceID)
	rows, err := stmt.Query(project, name)
	if err != nil {
		return -1, errors.Wrap(err, "Failed to get instance ID")
	}

	defer rows.Close()

	// Ensure we read one and only one row.
	if !rows.Next() {
		return -1, ErrNoSuchObject
	}
	var id int64
	err = rows.Scan(&id)
	if err != nil {
		return -1, errors.Wrap(err, "Failed to scan ID")
	}

	if rows.Next() {
		return -1, fmt.Errorf("More than one row returned")
	}
	err = rows.Err()
	if err != nil {
		return -1, errors.Wrap(err, "Result set failure")
	}

	return id, nil
}

// InstanceExists checks if a instance with the given key exists.
// generator: instance Exists
func (c *ClusterTx) InstanceExists(project string, name string) (bool, error) {
	_, err := c.GetInstanceID(project, name)
	if err != nil {
		if err == ErrNoSuchObject {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// CreateInstance adds a new instance to the database.
// generator: instance Create
func (c *ClusterTx) CreateInstance(object Instance) (int64, error) {
	// Check if a instance with the same key exists.
	exists, err := c.InstanceExists(object.Project, object.Name)
	if err != nil {
		return -1, errors.Wrap(err, "Failed to check for duplicates")
	}

	if exists {
		return -1, fmt.Errorf("This instance already exists")
	}

	args := make([]interface{}, 11)

	// Populate the statement arguments.
	args[0] = object.Project
	args[1] = object.Name
	args[2] = object.Node
	args[3] = object.Type
	args[4] = object.Architecture
	args[5] = object.Ephemeral
	args[6] = object.CreationDate
	args[7] = object.Stateful
	args[8] = object.LastUseDate
	args[9] = object.Description
	args[10] = object.ExpiryDate

	// Prepared statement to use.
	stmt := c.stmt(instanceCreate)

	// Execute the statement.
	result, err := stmt.Exec(args...)
	if err != nil {
		return -1, errors.Wrap(err, "Failed to create instance")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, errors.Wrap(err, "Failed to fetch instance ID")
	}

	return id, nil
}

// RenameInstance renames the instance matching the given key parameters.
// generator: instance Rename
func (c *ClusterTx) RenameInstance(project string, name string, to string) error {
	stmt := c.stmt(instanceRename)
	result, err := stmt.Exec(to, project, name)
	if err != nil {
		return errors.Wrap(err, "Rename instance")
	}

	n, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "Fetch affected rows")
	}

	if n != 1 {
		return fmt.Errorf("Query affected %d rows instead of 1", n)
	}
	return nil
}

// DeleteInstance deletes the instance matching the given key parameters.
// generator: instance DeleteOne-by-Project-and-Name
func (c *ClusterTx) DeleteInstance(project string, name string) error {
	stmt := c.stmt(instanceDeleteByProjectAndName)
	result, err := stmt.Exec(project, name)
	if err != nil {
		return errors.Wrap(err, "Delete instance")
	}

	n, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "Fetch affected rows")
	}

	if n != 1 {
		return fmt.Errorf("Query deleted %d rows instead of 1", n)
	}

	return nil
}

// UpdateInstance updates the instance matching the given key parameters.
// generator: instance Update
func (c *ClusterTx) UpdateInstance(project string, name string, object Instance) error {
	id, err := c.GetInstanceID(project, name)
	if err != nil {
		return errors.Wrap(err, "Get instance")
	}

	stmt := c.stmt(instanceUpdate)
	result, err := stmt.Exec(object.Project, object.Name, object.Node, object.Type, object.Architecture, object.Ephemeral, object.CreationDate, object.Stateful, object.LastUseDate, object.Description, object.ExpiryDate, id)
	if err != nil {
		return errors.Wrap(err, "Update instance")
	}

	n, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "Fetch affected rows")
	}

	if n != 1 {
		return fmt.Errorf("Query updated %d rows instead of 1", n)
	}

	return nil
}
