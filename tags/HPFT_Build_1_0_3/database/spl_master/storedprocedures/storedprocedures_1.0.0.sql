CREATE PROCEDURE sp_mst_chk_user_login(
	IN `in_usr_name` VARCHAR(150),
	IN `in_usr_password` VARCHAR(150)

)
COMMENT ''
BEGIN

SELECT * FROM spl_master_user_tbl 
WHERE usr_name = in_usr_name AND usr_password = in_usr_password AND usr_state = 1;

END;



CREATE PROCEDURE sp_mst_get_usr_products(
	IN `in_user_id` INT
)
COMMENT 'This procedure will get the associated products with login user'
BEGIN

select cpm.id as cpm_id,prod.prod_code,cpm.cust_id_fk,cust.cust_name from  spl_master_usr_cpm_tbl as ucpm 
inner join spl_master_cust_prod_mapping_tbl as cpm on ucpm.cpm_id_fk = cpm.id
inner join spl_master_product_tbl as prod on prod.id = cpm.prod_id_fk
inner join spl_master_customer_tbl as cust on cust.id = cpm.cust_id_fk
where cpm.cpm_state = 1 
AND cust.cust_state = 1
AND  ucpm.user_id_fk = in_user_id;

END;


CREATE PROCEDURE sp_mst_get_configuration()
COMMENT 'This procedure will get all the configuration'
BEGIN

SELECT * FROM  spl_master_config;

END;
