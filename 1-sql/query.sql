SELECT u.id as "ID", u.user_name as "UserName", pu.user_name as "ParentUserName" FROM users u LEFT JOIN users pu ON u.parent = pu.id;
