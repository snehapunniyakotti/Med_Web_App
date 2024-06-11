<template>
    <v-container >
    <v-card class="mx-auto" max-width="400" elevation="5">
        <v-card-title class="justify-center">
            LogIn <v-icon>mdi-login</v-icon>
        </v-card-title>
        <v-card-text>
            <!-- <v-form @submit.prevent="login"> -->
            <v-text-field v-model="userId" label="User ID" outlined></v-text-field>
            <v-text-field v-model="password" :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
                :type="showPassword ? 'text' : 'password'" label="Password" outlined
                @click:append="showPassword = !showPassword">
            </v-text-field>
            <!-- <v-text-field v-model="password" label="Password" outlined type="password"></v-text-field> -->
            <v-btn color="#81D4FA" class="white--text" type="submit" block @click="login">Login</v-btn>
            <v-snackbar v-model="snackbar" :timeout=2000 centered color="red"> {{ msg }}</v-snackbar>
            <!-- <v-alert class="mt-7" type="error" v-show="msg">{{ msg }}</v-alert>  -->
        </v-card-text>
    </v-card>
    </v-container>
</template>
<script>
import EventServices from "../services/EventServices";
export default {
    name: 'login',
    data() {
        return {
            userId: '',
            password: '',
            msg: '',
            showPassword: false,
            snackbar: false,
        };
    },
    methods: {
        login() {
            // checking that both fields data entered or not
            if (this.userId && this.password) {
                // here calling this login() present in EventServices.js and passsing user entered userId and password as parameter
                EventServices.login(this.userId, this.password)
                    // .then will handle response get from EventServices.login() 
                    .then(response => {
                        const data = response.data;
                        // Here checking the response status get from EventServices.js
                        if (data.resp.status === 'S') {
                            // pushing to dashboard based on role get as response from EventServices.login() method
                            this.$router.push({ name: 'dashboard', params: { role: data.role, Id: this.userId, lId: data.loginHistoryID } });
                        } else {
                            // console.log(data.errMsg)
                            this.snackbar = true
                            this.msg = data.errMsg || 'Login failed :  UserId and Password Not found';
                        }
                    })
                    //.catch will handle error. getting error other than response like axios..
                    .catch(error => {
                        // console.log(error)
                        this.snackbar = true
                        this.msg = 'An error occurred: Internal Error' + error.message;
                    });
            } else {
                this.snackbar = true
                this.msg = 'Incorrect Input, Please Check';
            }
        }
    }
}
</script>
