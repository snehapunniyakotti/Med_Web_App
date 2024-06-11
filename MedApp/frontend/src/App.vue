<template>
  <v-app >
  <!-- <v-container fluid fill-height :style="{ background: 'linear-gradient(to bottom, #81D4FA, #FFFFFF, #FFFFFF)', height: '100vh', padding: '20px', 'background-attachment': 'fixed','overflow': 'auto' }"> -->
    <v-img src="../src/assets/medappBg.png">
    <Navbar :type="istype" v-show="islogin" :id="id" :lid="lid" />
    <v-main class="">
      <router-view />

    </v-main>
  </v-img>
<!-- </v-container> -->
</v-app>
</template>

<script>
import Navbar from './components/NavComp/Navbar.vue';

export default {
  name: "App",

  data() {
    return {

    }
  },
  components: {
    Navbar,
  },
  computed: {

    islogin() {
      // console.log(this.$route.path)
      return (this.$route.path !== "/");
    },
    istype() {
      // console.log(this.$route.params.role +" from app.vue")
      return this.$route.params.role;
    },
    id(){
      return this.$route.params.Id;
    },
    lid(){
      return this.$route.params.lId;
    }
  },
  created() {

    if(this.$route.path !== "/" && this.$route.path !='/Errormsg'){
      let role = this.$route.params.role;
      if (role) {
        // console.log('hi')
      } else {
        this.$router.push('/Errormsg');
      }
    }
  },
  beforeUpdate(){
    if(this.$route.path !== "/" && this.$route.path !='/Errormsg'){
      let role = this.$route.params.role;
      if (role) {
        // console.log('hi')
      } else {
        this.$router.push('/Errormsg');
      }
    }
  }
};
</script>
