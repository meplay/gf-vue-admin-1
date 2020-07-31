import service from '@/utils/request'

// @Tags CasbinRule
// @Summary 创建CasbinRule
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CasbinRule true "创建CasbinRule"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /casbinRule/createCasbinRule [post]
export const createCasbinRule = (data) => {
     return service({
         url: "/casbinRule/createCasbinRule",
         method: 'post',
         data
     })
 }


// @Tags CasbinRule
// @Summary 删除CasbinRule
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CasbinRule true "删除CasbinRule"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /casbinRule/deleteCasbinRule [delete]
 export const deleteCasbinRule = (data) => {
     return service({
         url: "/casbinRule/deleteCasbinRule",
         method: 'delete',
         data
     })
 }

// @Tags CasbinRule
// @Summary 删除CasbinRule
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CasbinRule"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /casbinRule/deleteCasbinRule [delete]
 export const deleteCasbinRuleByIds = (data) => {
     return service({
         url: "/casbinRule/deleteCasbinRuleByIds",
         method: 'delete',
         data
     })
 }

// @Tags CasbinRule
// @Summary 更新CasbinRule
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CasbinRule true "更新CasbinRule"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /casbinRule/updateCasbinRule [put]
 export const updateCasbinRule = (data) => {
     return service({
         url: "/casbinRule/updateCasbinRule",
         method: 'put',
         data
     })
 }


// @Tags CasbinRule
// @Summary 用id查询CasbinRule
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CasbinRule true "用id查询CasbinRule"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /casbinRule/findCasbinRule [get]
 export const findCasbinRule = (params) => {
     return service({
         url: "/casbinRule/findCasbinRule",
         method: 'get',
         params
     })
 }


// @Tags CasbinRule
// @Summary 分页获取CasbinRule列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取CasbinRule列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /casbinRule/getCasbinRuleList [get]
 export const getCasbinRuleList = (params) => {
     return service({
         url: "/casbinRule/getCasbinRuleList",
         method: 'get',
         params
     })
 }