CREATE PROCEDURE `sp_mst_chk_user_login`(
	IN `in_usr_name` VARCHAR(150),
	IN `in_usr_password` VARCHAR(150)

)
COMMENT ''
BEGIN

SELECT * FROM spl_master_user_tbl 
WHERE usr_name = in_usr_name AND usr_password = in_usr_password AND usr_state = 1;

END