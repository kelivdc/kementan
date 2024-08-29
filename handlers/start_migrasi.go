package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"kementan.com/database"
)

func NullToZero(input interface{}) interface{} {
	if input == nil {
		return 0
	}
	return input
}

func StartMigrasi(c *fiber.Ctx) error {
	var count int64
	result := []map[string]interface{}{}
	sql := `SELECT * FROM submissions_simluhtan order by id ASC LIMIT 10`
	database.DB.Db.Raw(sql).Scan(&result)

	database.DB.Db.Raw("SELECT COUNT(id) as total FROM submissions_simluhtan_alo_mitreka3").Count(&count)

	for _, data := range result {
		if count > 0 {
			// UPDATE DATE
			fmt.Println("== Update ==")
		} else {
			// INSERT NEW
			sql_insert := `
			 INSERT INTO submissions_siluhtan_delta_2 (
			  year, sub_district_code,
			  status_old, status_new,
			  retailer_id_old, retailer_id_new,
			  farmer_group_id, farmer_nik,
			  mt1_commodity, mt2_commodity, mt3_commodity,
			  mt1_urea_old, mt1_urea_new, mt1_urea_delta,
			  mt2_urea_old, mt2_urea_new, mt2_urea_delta,
			  mt3_urea_old, mt3_urea_new, mt3_urea_delta,
			  mt1_sp36_old, mt1_sp36_new, mt1_sp36_delta,
			  mt2_sp36_old, mt2_sp36_new, mt2_sp36_delta,
			  mt3_sp36_old, mt3_sp36_new, mt3_sp36_delta,
			  mt1_za_old, mt1_za_new, mt1_za_delta,
			  mt2_za_old, mt2_za_new, mt2_za_delta,
			  mt3_za_old, mt3_za_new, mt3_za_delta,
			  mt1_npk_old, mt1_npk_new, mt1_npk_delta,
			  mt2_npk_old, mt2_npk_new, mt2_npk_delta,
			  mt3_npk_old, mt3_npk_new, mt3_npk_delta,
			  mt1_organic_old, mt1_organic_new, mt1_organic_delta,
			  mt2_organic_old, mt2_organic_new, mt2_organic_delta,
			  mt3_organic_old, mt3_organic_new, mt3_organic_delta,
			  mt1_npk_formula_old, mt1_npk_formula_new, mt1_npk_formula_delta,
			  mt2_npk_formula_old, mt2_npk_formula_new, mt2_npk_formula_delta,
			  mt3_npk_formula_old, mt3_npk_formula_new, mt3_npk_formula_delta,
			  mt1_poc_old, mt1_poc_new, mt1_poc_delta,
			  mt2_poc_old, mt2_poc_new, mt2_poc_delta,
			  mt3_poc_old, mt3_poc_new, mt3_poc_delta,
			  created_at, updated_at, id_agt, is_new, id_submission
			)
			VALUES
			  (
				2024, ?,
				?, ?,
				?, ?,
				?, ?,
				'', '', '',
				0, ?, ?,
				0, ?, ?,
				0, ?, ?,
				0, ?, ?,
				0, ?, ?,
				0, ?, ?,
				0, ?, ?,
				0, ?, ?,
				0, ?, ?,
				0, ?, ?,
				0, ?, ?,
				0, ?, ?,
				0, ?, ?,
				0, ?, ?,
				0, ?, ?,
				0, ?, ?,
				0, ?, ?,
				0, ?, ?,
				0, ?, ?,
				0, ?, ?,
				0, ?, ?,
				CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 0, 0, ?
			  );
			`
			database.DB.Db.Exec(sql_insert,
				data["sub_district_code"],
				NullToZero(data["var_status_old"]), NullToZero(data["var_status_new"]),
				NullToZero(data["var_retailer_id_old"]), NullToZero(data["var_retailer_id_new"]),
				NullToZero(data["farmer_group_id"]), NullToZero(data["farmer_nik"]),
				NullToZero(data["mt1_urea_new"]), NullToZero(data["mt1_urea_delta"]),
				NullToZero(data["mt2_urea_new"]), NullToZero(data["mt2_urea_delta"]),
				NullToZero(data["mt3_urea_new"]), NullToZero(data["mt3_urea_delta"]),
				NullToZero(data["mt1_sp36_new"]), NullToZero(data["mt1_sp36_delta"]),
				NullToZero(data["mt2_sp36_new"]), NullToZero(data["mt2_sp36_delta"]),
				NullToZero(data["mt3_sp36_new"]), NullToZero(data["mt3_sp36_delta"]),
				NullToZero(data["mt1_za_new"]), NullToZero(data["mt1_za_delta"]),
				NullToZero(data["mt2_za_new"]), NullToZero(data["mt2_za_delta"]),
				NullToZero(data["mt3_za_new"]), NullToZero(data["mt3_za_delta"]),
				NullToZero(data["mt1_npk_new"]), NullToZero(data["mt1_npk_delta"]),
				NullToZero(data["mt2_npk_new"]), NullToZero(data["mt2_npk_delta"]),
				NullToZero(data["mt3_npk_new"]), NullToZero(data["mt3_npk_delta"]),
				NullToZero(data["mt1_organic_new"]), NullToZero(data["mt1_organic_delta"]),
				NullToZero(data["mt2_organic_new"]), NullToZero(data["mt2_organic_delta"]),
				NullToZero(data["mt3_organic_new"]), NullToZero(data["mt3_organic_delta"]),
				NullToZero(data["mt1_npk_formula_new"]), NullToZero(data["mt1_npk_formula_delta"]),
				NullToZero(data["mt2_npk_formula_new"]), NullToZero(data["mt2_npk_formula_delta"]),
				NullToZero(data["mt3_npk_formula_new"]), NullToZero(data["mt3_npk_formula_delta"]),
				NullToZero(data["mt1_poc_new"]), NullToZero(data["mt1_poc_delta"]),
				NullToZero(data["mt2_poc_new"]), NullToZero(data["mt2_poc_delta"]),
				NullToZero(data["mt3_poc_new"]), NullToZero(data["mt3_poc_delta"]),
				NullToZero(data["id_old"]),
			)

			sql_alo := `
				INSERT INTO submissions_simluhtan_alo_mitreka3 (
				id, year, sub_district_code, district_code, 
				city_code, province_code, status, 
				rejection_message, instructor_name, 
				retailer_id, farmer_group_id, farmer_group_union_name, 
				farmer_name, farmer_nik, farmer_address, 
				farmer_mother_name, farmer_born_place, 
				farmer_born_date, subsector, mt1_planting_area, 
				mt2_planting_area, mt3_planting_area, 
				mt1_commodity, mt2_commodity, mt3_commodity, 
				mt1_urea, mt2_urea, mt3_urea, mt1_sp36, 
				mt2_sp36, mt3_sp36, mt1_za, mt2_za, 
				mt3_za, mt1_npk, mt2_npk, mt3_npk, 
				mt1_organic, mt2_organic, mt3_organic, 
				mt1_npk_formula, mt2_npk_formula, 
				mt3_npk_formula, mt1_poc, mt2_poc, 
				mt3_poc, field_coordinate, created_at, 
				updated_at, farmer_group_name, farmer_phone_number, 
				id_agt
				) 
				VALUES (
				?, ?, ?, ?, 
				?, ?, ?, 
				?, ?, 
				?, ?, ?, 
				?, ?, ?, 
				?, ?, 
				?, ?, ?, 
				?, ?, 
				?, ?, ?, 
				?, ?, ?, ?,
				?, ?, ?, ?, 
				?, ?, ?, ?, 
				?, ?, ?, 
				?, ?, 
				?, ?, ?,
				?, ?, ?,
				?, ?, ?, 
				?
				)
			`
			database.DB.Db.Exec(sql_alo,
				data["id"], data["year"], data["sub_district_code"], data["district_code"],
				data["city_code"], data["province_code"], data["status"],
				data["rejection_message"], data["instructor_name"],
				data["retailer_id"], data["farmer_group_id"], data["farmer_group_union_name"],
				data["farmer_name"], data["farmer_nik"], data["farmer_address"],
				data["farmer_mother_name"], data["farmer_born_place"],
				data["farmer_born_date"], data["subsector"], data["mt1_planting_area"],
				data["mt2_planting_area"], data["mt3_planting_area"],
				data["mt1_commodity"], data["mt2_commodity"], data["mt3_commodity"],
				data["mt1_urea"], data["mt2_urea"], data["mt3_urea"], data["mt1_sp36"],
				data["mt2_sp36"], data["mt3_sp36"], data["mt1_za"], data["mt2_za"],
				data["mt3_za"], data["mt1_npk"], data["mt2_npk"], data["mt3_npk"],
				data["mt1_organic"], data["mt2_organic"], data["mt3_organic"],
				data["mt1_npk_formula"], data["mt2_npk_formula"],
				data["mt3_npk_formula"], data["mt1_poc"], data["mt2_poc"],
				data["mt3_poc"], data["field_coordinate"], data["created_at"],
				data["updated_at"], data["farmer_group_name"], data["farmer_phone_number"],
				data["id_agt"],
			)
		}
	}
	return c.JSON(fiber.Map{
		"data": "Selesai",
	})
}

func District(c *fiber.Ctx) error {
	var total int64
	database.DB.Db.Distinct("district_code").Table("submissions_simluhtan").Count(&total)
	return c.JSON(fiber.Map{
		"district": total,
	})
}
