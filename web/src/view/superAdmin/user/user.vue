<template>
  <div>
    <div class="button-box clearflex">
      <el-button @click="addUser" type="primary">新增用户</el-button>
    </div>
    <el-table :data="tableData" border stripe>
      <el-table-column label="头像" min-width="50">
        <template slot-scope="scope">
          <div :style="{'textAlign':'center'}">
            <img :src="scope.row.header_img" height="35" width="35" />
          </div>
        </template>
      </el-table-column>
      <el-table-column label="uuid" min-width="250" prop="uuid"></el-table-column>
      <el-table-column label="用户名" min-width="150" prop="username"></el-table-column>
      <el-table-column label="昵称" min-width="150" prop="nickname"></el-table-column>
      <el-table-column label="用户角色" min-width="150">
        <template slot-scope="scope">
          <el-cascader
            @change="changeAuthority(scope.row)"
            v-model="scope.row.authority.authority_id"
            :options="authOptions"
            :show-all-levels="false"
            :props="{ checkStrictly: true,label:'authority_name',value:'authority_id',disabled:'disabled',emitPath:false}"
            filterable
          ></el-cascader>
        </template>
      </el-table-column>
      <el-table-column label="操作" min-width="150">
        <template slot-scope="scope">
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除此用户吗</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteUser(scope.row)">确定</el-button>
            </div>
            <el-button type="danger" icon="el-icon-delete" size="small" slot="reference">删除</el-button>
          </el-popover>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog :visible.sync="addUserDialog" custom-class="user-dialog" title="新增用户">
      <el-form :rules="rules" ref="userForm" :model="userInfo">
        <el-form-item label="用户名" label-width="80px" prop="username">
          <el-input v-model="userInfo.username"></el-input>
        </el-form-item>
        <el-form-item label="密码" label-width="80px" prop="password">
          <el-input v-model="userInfo.password"></el-input>
        </el-form-item>
        <el-form-item label="别名" label-width="80px" prop="nickName">
          <el-input v-model="userInfo.nickname"></el-input>
        </el-form-item>
        <el-form-item label="头像" label-width="80px">
          <el-upload
            :headers="{'x-token':token}"
            :on-success="handleAvatarSuccess"
            :show-file-list="false"
            :action="`${path}/fileUploadAndDownload/upload?noSave=1`"
            class="avatar-uploader"
            name="file"
          >
            <img :src="userInfo.header_img" class="avatar" v-if="userInfo.header_img" />
            <i class="el-icon-plus avatar-uploader-icon" v-else></i>
          </el-upload>
        </el-form-item>
        <el-form-item label="用户角色" label-width="80px" prop="authority_id">
          <el-cascader
            @change="changeAuthority(scope.row)"
            v-model="userInfo.authority_id"
            :options="authOptions"
            :show-all-levels="false"
            :props="{ checkStrictly: true,label:'authority_name',value:'authority_id',disabled:'disabled',emitPath:false}"
            filterable
          ></el-cascader>
        </el-form-item>
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeAddUserDialog">取 消</el-button>
        <el-button @click="enterAddUserDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>


<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成
const path = process.env.VUE_APP_BASE_API;
import {
  getUserList,
  setUserAuthority,
  register,
  deleteUser
} from "@/api/user";
import { getAuthorityList } from "@/api/authority";
import infoList from "@/components/mixins/infoList";
import { mapGetters } from "vuex";
export default {
  name: "Api",
  mixins: [infoList],
  data() {
    return {
      listApi: getUserList,
      path: path,
      authOptions: [],
      addUserDialog: false,
      userInfo: {
        username: "",
        password: "",
       nickname: "",
        header_img: "",
        authority_id: ""
      },
      rules: {
        username: [
          { required: true, message: "请输入用户名", trigger: "blur" },
          { min: 6, message: "最低6位字符", trigger: "blur"}
        ],
        password: [
          { required: true, message: "请输入用户密码", trigger: "blur" },
          { min: 6, message: "最低6位字符", trigger: "blur"}
        ],
       nickname: [
          { required: true, message: "请输入用户昵称", trigger: "blur" }
        ],
        authority_id: [
          { required: true, message: "请选择用户角色", trigger: "blur" }
        ]
      }
    };
  },
  computed: {
    ...mapGetters("user", ["token"])
  },
  methods: {
    setOptions(authData) {
      this.authOptions = [];
      this.setAuthorityOptions(authData, this.authOptions);
    },
    setAuthorityOptions(AuthorityData, optionsData) {
      AuthorityData &&
        AuthorityData.map(item => {
          if (item.children&&item.children.length) {
            const option = {
              authority_id: item.authority_id,
              authority_name: item.authority_name,
              children: []
            };
            this.setAuthorityOptions(item.children, option.children);
            optionsData.push(option);
          } else {
            const option = {
              authority_id: item.authority_id,
              authority_name: item.authority_name
            };
            optionsData.push(option);
          }
        });
    },
    async deleteUser(row) {
      const res = await deleteUser({ id: row.id });
      if (res.code == 0) {
        this.getTableData();
        row.visible = false;
      }
    },
    async enterAddUserDialog() {
      this.$refs.userForm.validate(async valid => {
        if (valid) {
          const res = await register(this.userInfo);
          if (res.code == 0) {
            this.$message({ type: "success", message: "创建成功" });
          }
          await this.getTableData();
          this.closeAddUserDialog();
        }
      });
    },
    closeAddUserDialog() {
      this.$refs.userForm.resetFields();
      this.addUserDialog = false;
    },
    handleAvatarSuccess(res) {
      this.userInfo.header_img = res.data.file.url;
    },
    addUser() {
      this.addUserDialog = true;
    },
    async changeAuthority(row) {
      const res = await setUserAuthority({
        uuid: row.uuid,
        authority_id: row.authority.authority_id
      });
      if (res.code == 0) {
        this.$message({ type: "success", message: "角色设置成功" });
      }
    }
  },
  async created() {
    this.getTableData();
    const res = await getAuthorityList({ page: 1, pageSize: 999 });
    this.setOptions(res.data.list);
  }
};
</script>
<style scoped lang="scss">
.button-box {
  padding: 10px 20px;
  .el-button {
    float: right;
  }
}

.user-dialog {
  .avatar-uploader .el-upload:hover {
    border-color: #409eff;
  }
  .avatar-uploader-icon {
    border: 1px dashed #d9d9d9 !important;
    border-radius: 6px;
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    line-height: 178px;
    text-align: center;
  }
  .avatar {
    width: 178px;
    height: 178px;
    display: block;
  }
}
</style>