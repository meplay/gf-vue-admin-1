<template>
  <div class="system">
    <el-form :model="config" label-width="100px" ref="form" class="system">
      <h1>仅支持查看服务器的配置文件,不支持修改与重启服务器</h1>
      <!-- jwt签名 start -->
      <h2>jwt签名</h2>
      <el-form-item label="jwt签名">
        <el-input v-model="config.jwt.signing_key"></el-input>
      </el-form-item>
      <el-form-item label="过期时间">
        <el-input v-model.number="config.jwt.expires_at"></el-input>
      </el-form-item>
      <el-form-item label="刷新时间">
        <el-input v-model.number="config.jwt.refresh_at"></el-input>
      </el-form-item>
      <!-- jwt签名 end -->

      <!-- Oss start -->
      <template v-if="config.system.oss_type == 'local'">
        <h2>local配置</h2>
        <el-form-item label="local_path">
          <el-input v-model="config.oss.local.local_path"></el-input>
        </el-form-item>
      </template>
      <template v-if="config.system.oss_type == 'qiniu'">
        <h2>qiniu配置</h2>
        <el-form-item label="zone">
          <el-input v-model="config.oss.qiniu.zone"></el-input>
        </el-form-item>
        <el-form-item label="bucket">
          <el-input v-model="config.oss.qiniu.bucket"></el-input>
        </el-form-item>
        <el-form-item label="ImgPath">
          <el-input v-model="config.oss.qiniu.img_path"></el-input>
        </el-form-item>
        <el-form-item label="UseHttps">
          <el-checkbox v-model="config.oss.qiniu.use_https"></el-checkbox>
        </el-form-item>
        <el-form-item label="SecretKey">
          <el-input v-model="config.oss.qiniu.secret_key"></el-input>
        </el-form-item>
        <el-form-item label="UseCdnDomains">
          <el-checkbox v-model="config.oss.qiniu.use_cdn_domains"></el-checkbox>
        </el-form-item>
      </template>
      <template v-if="config.system.oss_type == 'minio'">
        <h2>minio配置</h2>
        <el-form-item label="Id">
          <el-input v-model="config.oss.minio.id"></el-input>
        </el-form-item>
        <el-form-item label="Path">
          <el-input v-model="config.oss.minio.path"></el-input>
        </el-form-item>
        <el-form-item label="Token">
          <el-input v-model="config.oss.minio.token"></el-input>
        </el-form-item>
        <el-form-item label="Bucket">
          <el-input v-model="config.oss.minio.bucket"></el-input>
        </el-form-item>
        <el-form-item label="UseSsl">
          <el-checkbox v-model="config.oss.minio.use_ssl"></el-checkbox>
        </el-form-item>
        <el-form-item label="Secret">
          <el-input v-model="config.oss.minio.secret"></el-input>
        </el-form-item>
        <el-form-item label="Endpoint">
          <el-input v-model="config.oss.minio.endpoint"></el-input>
        </el-form-item>
      </template>
      <template v-if="config.system.oss_type == 'aliyun'">
        <h2>Aliyun配置</h2>
        <el-form-item label="Path">
          <el-input v-model="config.oss.aliyun.path"></el-input>
        </el-form-item>
        <el-form-item label="Bucket">
          <el-input v-model="config.oss.aliyun.bucket"></el-input>
        </el-form-item>
        <el-form-item label="ACLType">
          <el-input v-model="config.oss.aliyun.acl_type"></el-input>
        </el-form-item>
        <el-form-item label="Endpoint">
          <el-input v-model="config.oss.aliyun.endpoint"></el-input>
        </el-form-item>
        <el-form-item label="AccessKeyID">
          <el-input v-model="config.oss.aliyun.access_key_id"></el-input>
        </el-form-item>
        <el-form-item label="SecretAccessKey">
          <el-input v-model="config.oss.aliyun.secret_access_key"></el-input>
        </el-form-item>
        <el-form-item label="StorageClassType">
          <el-input v-model="config.oss.aliyun.storage_class_type"></el-input>
        </el-form-item>
      </template>
      <!-- Oss end -->

      <!-- Email start -->
      <h2>Redis数据库配置</h2>
      <el-form-item label="Default">
        <el-input v-model="config.redis.default"></el-input>
      </el-form-item>
      <el-form-item label="Cache">
        <el-input v-model="config.redis.cache"></el-input>
      </el-form-item>
      <!-- Oss end -->

      <!-- Email start -->
      <h2>邮箱配置</h2>
      <el-form-item label="To">
        <el-input v-model="config.email.to"></el-input>
      </el-form-item>
      <el-form-item label="Port">
        <el-input v-model.number="config.email.port"></el-input>
      </el-form-item>
      <el-form-item label="From">
        <el-input v-model="config.email.from" placeholder="可多个，以逗号分隔(单个时请不要加)"></el-input>
      </el-form-item>
      <el-form-item label="Host">
        <el-input v-model="config.email.host"></el-input>
      </el-form-item>
      <el-form-item label="IsSsl">
        <el-checkbox v-model="config.email.is_ssl"></el-checkbox>
      </el-form-item>
      <el-form-item label="Secret">
        <el-input v-model="config.email.secret"></el-input>
      </el-form-item>
      <el-form-item label="Nickname">
        <el-input v-model="config.email.nickname"></el-input>
      </el-form-item>
      <el-form-item label="测试邮件">
        <el-button @click="email">测试邮件</el-button>
      </el-form-item>
      <!-- Email end -->

      <!-- Casbin start -->
      <h2>casbin配置</h2>
      <el-form-item label="模型地址">
        <el-input v-model="config.casbin.model_path"></el-input>
      </el-form-item>
      <!-- Casbin end -->

      <!-- Logger start -->
      <h2>Logger配置</h2>
      <el-form-item label="Path">
        <el-input v-model="config.logger.path"></el-input>
      </el-form-item>
      <el-form-item label="Level">
        <el-input v-model="config.logger.level"></el-input>
      </el-form-item>
      <el-form-item label="Stdout">
        <el-checkbox v-model="config.logger.stdout"></el-checkbox>
      </el-form-item>
      <!-- Logger end -->

      <!-- Server start -->
      <h2>Server配置</h2>
      <el-form-item label="LogPath">
        <el-input v-model="config.server.log_path"></el-input>
      </el-form-item>
      <el-form-item label="Address">
        <el-input v-model="config.server.address"></el-input>
      </el-form-item>
      <h3>指定服务器启动时是否自动转储路由器映射</h3>
      <el-form-item label="是否开启">
        <el-checkbox v-model="config.server.dump_router_map"></el-checkbox>
      </el-form-item>
      <h3>启用将内容记录到文件的错误</h3>
      <el-form-item label="是否开启">
        <el-checkbox v-model="config.server.error_log_enabled"></el-checkbox>
      </el-form-item>
      <h3>启用对文件的日志记录内容访问</h3>
      <el-form-item label="是否开启">
        <el-checkbox v-model="config.server.access_log_enabled"></el-checkbox>
      </el-form-item>
      <!-- Server end -->

      <!-- System start -->
      <h2>系统配置</h2>
      <el-form-item label="db">
        <el-input v-model="config.system.db"></el-input>
      </el-form-item>
      <el-form-item label="环境值">
        <el-input v-model="config.system.env"></el-input>
      </el-form-item>
      <el-form-item label="OSS类别">
        <el-select v-model="config.system.oss_type">
          <el-option value="local"></el-option>
          <el-option value="qiniu"></el-option>
          <el-option value="minio"></el-option>
          <el-option value="aliyun"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="ErrorToEmail">
        <el-checkbox v-model="config.system.error_to_email">开启</el-checkbox>
      </el-form-item>
      <el-form-item label="多点登录拦截">
        <el-checkbox v-model="config.system.use_multipoint">开启</el-checkbox>
      </el-form-item>
      <!-- System end -->

      <!-- Captcha start -->
      <h2>验证码配置</h2>
      <el-form-item label="keyLong">
        <el-input v-model.number="config.captcha.key_long"></el-input>
      </el-form-item>
      <el-form-item label="imgWidth">
        <el-input v-model.number="config.captcha.img_width"></el-input>
      </el-form-item>
      <el-form-item label="imgHeight">
        <el-input v-model.number="config.captcha.img_height"></el-input>
      </el-form-item>
      <!-- Captcha end -->

      <!-- Database start -->
      <h2>Database</h2>
      <el-form-item label="Host">
        <el-input v-model="config.database.host"></el-input>
      </el-form-item>
      <el-form-item label="Port">
        <el-input v-model="config.database.user"></el-input>
      </el-form-item>
      <el-form-item label="User">
        <el-input v-model="config.database.pass"></el-input>
      </el-form-item>
      <el-form-item label="Type">
        <el-input v-model="config.database.type"></el-input>
      </el-form-item>
      <el-form-item label="Role">
        <el-input v-model="config.database.role"></el-input>
      </el-form-item>
      <el-form-item label="Debug">
        <el-checkbox v-model="config.database.debug"></el-checkbox>
      </el-form-item>
      <el-form-item label="Prefix">
        <el-input v-model="config.database.prefix"></el-input>
      </el-form-item>
      <el-form-item label="DryRun">
        <el-checkbox v-model="config.database.dry_run"></el-checkbox>
      </el-form-item>
      <el-form-item label="Weight">
        <el-input v-model.number="config.database.weight"></el-input>
      </el-form-item>
      <el-form-item label="Charset">
        <el-input v-model="config.database.charset"></el-input>
      </el-form-item>
      <el-form-item label="LinkInfo">
        <el-input v-model="config.database.link_info"></el-input>
      </el-form-item>
      <el-form-item label="MaxIdleConnCount">
        <el-input v-model.number="config.database.max_idle_conn_count"></el-input>
      </el-form-item>
      <el-form-item label="MaxOpenConnCount">
        <el-input v-model.number="config.database.max_open_conn_count"></el-input>
      </el-form-item>
      <el-form-item label="MaxConnLifetime">
        <el-input v-model="config.database.max_conn_lifetime"></el-input>
      </el-form-item>

      <!-- Database end -->

      <!-- DatabaseLogger start -->
      <h2>DatabaseLogger</h2>
      <el-form-item label="Path">
        <el-input v-model="config.database_logger.path"></el-input>
      </el-form-item>
      <el-form-item label="Level">
        <el-input v-model="config.database_logger.level"></el-input>
      </el-form-item>
      <el-form-item label="Stdout">
        <el-checkbox v-model="config.database_logger.stdout"></el-checkbox>
      </el-form-item>
      <!-- DatabaseLogger end -->

<!--      <el-form-item>-->
<!--        <el-button @click="update" type="primary">立即更新</el-button>-->
<!--        <el-button @click="reload" type="primary">重启服务（开发中）</el-button>-->
<!--      </el-form-item>-->
    </el-form>
  </div>
</template>

<script>
import {getSystemConfig, setSystemConfig} from "@/api/system";
import {emailTest} from "@/api/email";

export default {
  name: "Config",
  data() {
    return {
      config: {
        jwt: {},
        oss: {
          local:{},
          qiniu:{},
          minio:{},
          aliyun:{},
        },
        redis: {},
        email: {},
        casbin: {},
        logger: {},
        server: {},
        system: {},
        captcha: {},
        database: {},
        database_logger: {},
      }
    };
  },
  async created() {
    await this.initForm();
  },
  methods: {
    async initForm() {
      const res = await getSystemConfig();
      if (res.code == 0) {
        this.config = res.data.config;
      }
    },
    reload() {
    },
    async update() {
      const res = await setSystemConfig({config: this.config});
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "配置文件设置成功"
        });
        await this.initForm();
      }
    },
    async email() {
      const res = await emailTest();
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "邮件发送成功"
        });
        await this.initForm();
      } else {
        this.$message({
          type: "error",
          message: "邮件发送失败"
        });
      }
    }
  }
};
</script>
<style lang="scss">
.system {

h2 {
  padding: 10px;
  margin: 10px 0;
  font-size: 16px;
  box-shadow: -4px 1px 3px 0px #e7e8e8;
}

}
</style>
