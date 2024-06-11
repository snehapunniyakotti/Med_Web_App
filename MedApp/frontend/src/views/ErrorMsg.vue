<template>
        <v-card class=" text-h4  white--text ma-16" color="#616161" >
            "An unexpected error occurred. Please try again later"
            <v-btn color="#000000" class="white--text" @click="goLogIn">Logout</v-btn>
        </v-card>
</template>
<script>
export default {
    name:'Errormsg',
    data(){
        return{
            // role:'None'
        }
    },
    methods:{
        goLogIn(){
            this.$router.push('/');
        },
    }
}
</script>