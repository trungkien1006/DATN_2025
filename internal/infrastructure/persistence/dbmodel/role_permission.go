package dbmodel

import rolepermission "final_project/internal/domain/role_permission"

type RolePermission struct {
	RoleID       uint `gorm:"primaryKey"`
	PermissionID uint `gorm:"primaryKey"`

	Role       Role       `gorm:"foreignKey:RoleID"`
	Permission Permission `gorm:"foreignKey:PermissionID"`
}

// Domain -> DB
func RolePerDomainToDB(domainRolePer rolepermission.RolePermission) RolePermission {
	return RolePermission{
		RoleID:       domainRolePer.RoleID,
		PermissionID: domainRolePer.PermissionID,
	}
}

// DB -> Domain
func RolePerDBToDomain(db RolePermission) rolepermission.RolePermission {
	return rolepermission.RolePermission{
		RoleID:       db.RoleID,
		PermissionID: db.PermissionID,
	}
	// Các trường CreatedAt, UpdatedAt, DeletedAt sẽ để GORM xử lý
}

// DB -> Domain
func DBToRolePermissions(db Role) rolepermission.RolePermissionList {
	permissons := make([]rolepermission.Permission, 0)

	for _, value := range db.RolePermissions {
		permissons = append(permissons, rolepermission.Permission{
			Code: value.Permission.Code,
		})
	}

	return rolepermission.RolePermissionList{
		ID:          db.ID,
		Name:        db.Name,
		Permissions: permissons,
	}
}
