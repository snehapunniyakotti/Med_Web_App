<template>
    <v-container>
        <!-- hi this is dashboard for :  {{ role }} -->
                <!-- <v-btn class="primary" elevation="2" outlined @click="logout">Log Out</v-btn> -->

        <MainDashBoard :role="role"></MainDashBoard>
    </v-container>
</template>
<script>
import MainDashBoard from '../components/DashBoard/MainDashBoard.vue';
export default {
    name:'Dashboard',
    data(){
        return{
            role:'None',
        }
    },
    created(){
        let role=this.$route.params.role;
            if(role){
                this.role= role
            }
            // else{
        //         this.$router.push('/Errormsg');
        //     }
    },
    components:{
        MainDashBoard,
    },
    
}
</script>