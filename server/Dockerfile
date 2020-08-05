FROM loads/alpine:3.8

LABEL maintainer="SliverHorn"

###############################################################################
#                                INSTALLATION
###############################################################################

# 设置固定的项目路径
WORKDIR /gf-vue-admin

# 添加应用可执行文件，并设置执行权限
ADD main   .
RUN chmod +x /gf-vue-admin/main

# 添加I18N多语言文件、静态文件、配置文件、模板文件
ADD i18n     /gf-vue-admin/i18n
ADD public   /gf-vue-admin/public
ADD config   /gf-vue-admin/config
ADD template /gf-vue-admin/template

RUN ls -al && pwd

EXPOSE 8888
###############################################################################
#                                   START
###############################################################################
ENTRYPOINT /gf-vue-admin/main
