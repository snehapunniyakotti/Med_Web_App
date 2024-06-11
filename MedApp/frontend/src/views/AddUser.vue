<template>
  <v-container  class="mt-12 pa-10 ">

    <v-form ref="form">
      <v-row>
        <h2>Add User</h2>
        <v-icon> mdi-account</v-icon>
      </v-row>

      <!-- Input Fields and Button -->
      <v-row>
        <!-- User ID Input -->
        <v-text-field v-model.lazy="userId" label="User ID" :rules="userIdRules" outlined></v-text-field>
      </v-row>
      <v-row>
        <!-- Password Input -->
        <v-text-field v-model.lazy="password" :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
                :type="showPassword ? 'text' : 'password'" label="Password" outlined :rules="passwordRules"
                @click:append="showPassword = !showPassword">
            </v-text-field>
        <!-- <v-text-field v-model.lazy="password" label="Password" type="password" :rules="passwordRules"
          outlined></v-text-field> -->
      </v-row>
      <v-row>
        <!-- Dropdown Box -->
        <v-select v-model="role" :items="roles" label="Role" :rules="roleRules" outlined></v-select>
      </v-row>
      <v-row>
        <!-- Add Button -->
        <v-btn color="#81D4FA" @click="addUser">Add</v-btn>

        <!-- <v-col cols="12">
                <div class="text-h4 red--text">{{ msg }}</div>
            </v-col> -->
      </v-row>
      <v-snackbar v-model="snackbar1" :timeout="timeout" centered color="error">
        {{ msg }}</v-snackbar>
      <v-snackbar v-model="snackbar2" :timeout="timeout" centered color="success">
        {{ msg }}</v-snackbar>
      <v-snackbar v-model="snackbar3" :timeout="timeout" centered color="error">
        Incorrect Input, Please Check</v-snackbar>
    </v-form>
  </v-container>
</template>
<script>
import EventServices from '../services/EventServices';
export default {
  name: "Add_User",
  data() {
    return {
      roles: ["Biller", "Manager", "SystemAdmin", "Inventry"],
      userId: "",
      password: "",
      role: "",
      msg: "",
      timeout: 2000,
      snackbar1: false,
      snackbar2: false,
      snackbar3: false,
      userIdRules: [
        (v) => !!v || "User ID is required",
        (v) => (v && v.length >= 4) || "User ID must be at least 4 characters",
        (v) =>
          /^[A-Z]{1}/.test(v) ||
          "First letter must be uppercase alphabets",
      ],
      passwordRules: [
        (v) => !!v || "Password is required",
        (v) => (v && v.length >= 8) || "Password must be at least 8 characters",
      ],
      roleRules: [
        (v) => !!v || "Role is required",
      ],
      showPassword: false,

    };
  },
  methods: {
    checkinguserid(v) {
      return !!v && v && v.length >= 4 && /^[A-Z]{1}/.test(v);
    },
    checkingpass(v) {
      return !!v && v && v.length >= 8;
    },
    addUser() {
      // this.$refs.t1.reset()
     

      if (this.checkinguserid(this.userId) && this.checkingpass(this.password) && this.role) {
        // Call the API using EventServices
        EventServices.postUser(this.userId, this.password, this.role)
          .then(response => {
            const data = response.data;
            if (data.resp.status === 'S') {
              this.msg = data.msg;
              this.snackbar2 = true;
            } else {
              this.msg = data.resp.errMsg;
              this.snackbar1 = true
            }
            this.$refs.form.reset();
            // this.userId = "";
            // this.password = "";
            // this.role = "";
          })
          .catch(error => {
            console.error('API Error:', error);
            this.$refs.form.reset();
            // this.userId = "";
            // this.password = "";
            // this.role = "";
          });
      } else {
        this.snackbar3 = true;
      }
    },
  },
  
  
};
</script>
