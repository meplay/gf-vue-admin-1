<template>
  <div class="system">
    <el-form :model="config" label-width="200px" ref="form" class="system">
      <!--  System start  -->
      <h2>系统配置</h2>
      <el-form-item label="环境值">
        <el-input v-model="config.system.env"></el-input>
      </el-form-item>
      <el-form-item label="Oss类型">
        <el-select v-model="config.system.ossType">
          <el-option value="local"></el-option>
          <el-option value="qiniu"></el-option>
          <el-option value="minio"></el-option>
          <el-option value="aliyun"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="错误发送邮箱">
        <el-checkbox v-model="config.system.errorToEmail">开启</el-checkbox>
      </el-form-item>
      <el-form-item label="多点登录拦截">
        <el-checkbox v-model="config.system.useMultipoint">开启</el-checkbox>
      </el-form-item>
      <!--  System end  -->

      <!--  JWT start  -->
      <h2>jwt签名</h2>
      <el-form-item label="jwt签名">
        <el-input v-model.number="config.jwt.expiresAt"></el-input>
        <el-input v-model.number="config.jwt.refreshAt"></el-input>
        <el-input v-model="config.jwt.signingKey"></el-input>
      </el-form-item>
      <!--  JWT end  -->

      <!--  Redis start  -->
      <h2>Redis admin数据库配置</h2>
      <el-form-item label="db">
        <el-input v-model="config.redis.db"></el-input>
      </el-form-item>
      <el-form-item label="address">
        <el-input v-model="config.redis.address"></el-input>
      </el-form-item>
      <el-form-item label="password">
        <el-input v-model="config.redis.password"></el-input>
      </el-form-item>
      <!--  Redis end  -->

      <!--  Email start  -->
      <h2>邮箱配置</h2>
      <el-form-item label="接收者邮箱">
        <el-input v-model="config.email.to" placeholder="可多个，以逗号分隔"></el-input>
      </el-form-item>
      <el-form-item label="端口">
        <el-input v-model.number="config.email.port"></el-input>
      </el-form-item>
      <el-form-item label="发送者邮箱">
        <el-input v-model="config.email.from"></el-input>
      </el-form-item>
      <el-form-item label="host">
        <el-input v-model="config.email.host"></el-input>
      </el-form-item>
      <el-form-item label="是否为ssl">
        <el-checkbox v-model="config.email.isSSL"></el-checkbox>
      </el-form-item>
      <el-form-item label="secret">
        <el-input v-model="config.email.secret"></el-input>
      </el-form-item>
      <el-form-item label="测试邮件">
        <el-button @click="email">测试邮件</el-button>
      </el-form-item>
      <!--  Email end  -->

      <!--  Casbin start  -->
      <h2>casbin配置</h2>
      <el-form-item label="模型地址">
        <el-input v-model="config.casbin.modelPath"></el-input>
      </el-form-item>
      <!--  Casbin end  -->

      <!--  Captcha start  -->
      <h2>验证码配置</h2>
      <el-form-item label="keyLong">
        <el-input v-model.number="config.captcha.keyLong"></el-input>
      </el-form-item>
      <el-form-item label="imageWidth">
        <el-input v-model.number="config.captcha.imageWidth"></el-input>
      </el-form-item>
      <el-form-item label="imageHeight">
        <el-input v-model.number="config.captcha.imageHeight"></el-input>
      </el-form-item>
      <el-form-item label="验证码存在redis">
      <el-checkbox v-model="config.system.captchaInRedis">开启</el-checkbox>
      </el-form-item>
      <!--  Captcha end  -->

      <!--  ossType start  -->
      <template v-if="config.system.ossType == 'local'">
        <h2>本地上传配置</h2>
        <el-form-item label="本地文件路径">
          <el-input v-model="config.local.path"></el-input>
        </el-form-item>
      </template>
      <template v-if="config.system.ossType == 'qiniu'">
        <h2>qiniu上传配置</h2>
        <el-form-item label="存储区域">
          <el-input v-model="config.qiniu.zone"></el-input>
        </el-form-item>
        <el-form-item label="空间名称">
          <el-input v-model="config.qiniu.bucket"></el-input>
        </el-form-item>
        <el-form-item label="CDN加速域名">
          <el-input v-model="config.qiniu.imgPath"></el-input>
        </el-form-item>
        <el-form-item label="是否使用https">
          <el-checkbox v-model="config.qiniu.useHttps">开启</el-checkbox>
        </el-form-item>
        <el-form-item label="accessKey">
          <el-input v-model="config.qiniu.accessKey"></el-input>
        </el-form-item>
        <el-form-item label="secretKey">
          <el-input v-model="config.qiniu.secretKey"></el-input>
        </el-form-item>
        <el-form-item label="上传是否使用CDN上传加速">
          <el-checkbox v-model="config.qiniu.useCdnDomains">开启</el-checkbox>
        </el-form-item>
      </template>
      <template v-if="config.system.ossType == 'minio'">
        <h2>minio上传配置</h2>
        <el-form-item label="存储区域">
          <el-input v-model="config.minio.id"></el-input>
        </el-form-item>
        <el-form-item label="path">
          <el-input v-model="config.minio.path"></el-input>
        </el-form-item>
        <el-form-item label="token">
          <el-input v-model="config.minio.token"></el-input>
        </el-form-item>
        <el-form-item label="空间名称">
          <el-input v-model="config.minio.bucket"></el-input>
        </el-form-item>
        <el-form-item label="上传是否使用Ssl">
          <el-checkbox v-model="config.minio.useSsl">开启</el-checkbox>
        </el-form-item>
        <el-form-item label="secret">
          <el-input v-model="config.minio.secret"></el-input>
        </el-form-item>
        <el-form-item label="endpoint">
          <el-input v-model="config.minio.endpoint"></el-input>
        </el-form-item>
      </template>
      <template v-if="config.system.ossType == 'aliyun'">
        <h2>aliyun上传配置</h2>
        <el-form-item label="path">
          <el-input v-model="config.aliyun.path"></el-input>
        </el-form-item>
        <el-form-item label="空间名称">
          <el-input v-model="config.aliyun.bucket"></el-input>
        </el-form-item>
        <el-form-item label="acl-type">
          <el-input v-model="config.aliyun.aclType"></el-input>
        </el-form-item>
        <el-form-item label="endpoint">
          <el-input v-model="config.aliyun.endpoint"></el-input>
        </el-form-item>
        <el-form-item label="access-key-id">
          <el-input v-model="config.aliyun.accessKeyId"></el-input>
        </el-form-item>
        <el-form-item label="secret-access-key">
          <el-input v-model="config.aliyun.secretAccessKey"></el-input>
        </el-form-item>
        <el-form-item label="storage-class-type">
          <el-input v-model="config.aliyun.storageClassType"></el-input>
        </el-form-item>
      </template>
      <!--  ossType end  -->

      <el-form-item>
        <el-button @click="update" type="primary">立即更新</el-button>
        <el-button @click="reload" type="primary">重启服务（开发中）</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { getSystemConfig, setSystemConfig } from "@/api/system";
import { emailTest } from "@/api/email";
export default {
  name: "Config",
  data() {
    return {
      config: {
        system: {},
        jwt: {},
        casbin: {},
        redis: {},
        local: {},
        qiniu: {},
        minio: {},
        aliyun: {},
        captcha: {},
        email: {}
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
    reload() {},
    async update() {
      const res = await setSystemConfig({ config: this.config });
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
