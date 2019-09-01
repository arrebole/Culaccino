<template>
  <div>
    <Header />
    <article>
      <section class="login-form">
        <form  @submit.prevent="login">
          <div class="login-header"><h4>Sign in</h4></div>

          <div class="login-body">
            <div class="form-group">
              <label>Username</label>
              <input
                type="text"
                class="form-control"
                v-model="userName"
                placeholder="Enter your username..."
                autocomplete="off"
              />
            </div>
            <div class="form-group">
              <label>Password</label>
              <input
                type="password"
                class="form-control"
                v-model="password"
                placeholder="Enter your Password..."
                autocomplete="off"
              />
            </div>
            <button type="submit" class="btn">
              <span class="btn-title">Sign in</span>
            </button>
          </div>
        </form>
      </section>
    </article>
    <Footer />
  </div>
</template>


<script lang="ts">
import Vue from "vue";
import api from "../api/index";
import Header from "../components/Header.vue";
import Footer from "../components/Footer.vue";
import IResp from "../types/resp";
import util from "../util";

export default Vue.extend({
  data() {
    return {
      userName: "",
      password: "",
      cookie: ""
    };
  },
  methods: {
    async login() {
      const isOk = await this.fetchLogin();
      if (isOk) this.handleCookie();
      else alert("登录失败");
    },
    async fetchLogin(): Promise<boolean> {
      const res: IResp = await api.sessionLogin(this.user());
      if (res.message != "success" || res.data.token == null) return false;
      this.cookie = res.data.token;
      return true;
    },
    user(): { userName: string; password: string } {
      return { userName: this.userName, password: this.password };
    },
    handleCookie() {
      if (util.getCookie() != this.cookie) util.setCookie(this.cookie);
      this.gotoAdmin();
    },
    gotoAdmin() {
      this.$router.replace({ name: 'Admin'})
    }
  },
  components: {
    Header,
    Footer
  }
});
</script>

<style lang="scss" scoped>
article {
  background: #eee;
  padding: 30px;
  min-height: 760px;
}
.login-form {
  background-color: #fff;
  padding: 10 15px;
  margin: 20px auto;
  max-width: 415px;
  width: 80%;
  border: 1px solid rgb(223, 223, 223);
}
.login-header {
  padding: 30px 20px;
  color: rgb(141, 141, 141);

  h4 {
    margin: 0;
  }
}

.login-body {
  box-sizing: border-box;
  border-top: 1px solid rgba(0, 0, 0, 0.1);
  padding: 20px 0;
  margin: 0 20px;
  height: 239px;
}
.form-control {
  display: block;
  font-size: 0.9rem;
  font-family: Roboto, "sans serif";
  font-weight: 300;
  width: 100%;
  padding: 10px 0px;
  color: #333;
  border: none;
  box-shadow: none;
  border: none;
  border-bottom: 1px solid #e6e6e6;
  outline: none;
}
.form-group {
  margin-bottom: 20px;
}
.form-group > label {
  color: #24a0d1;
  font-size: 0.9rem;
}
.btn {
  display: block;
  cursor: pointer;
  width: 100%;
  background-color: #1f1f1f;
  border-color: #1f1f1f;
  color: #fff;
  padding: 14px 40px;
  font-size: 1rem;
  &:hover {
    background-color: #414141;
  }
}
</style>
