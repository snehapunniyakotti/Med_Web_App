<template>
    <v-container >
  <!-- <v-parallax src="https://cdn.vuetifyjs.com/images/backgrounds/vbanner.jpg"> -->

        <ForManager v-if="role=='Manager'" ></ForManager>
        <ForNormal v-if="role=='Inventry'|| role== 'SystemAdmin'"></ForNormal>
        <ForBiller v-if="role=='Biller'" ></ForBiller>
    <!-- </v-parallax> -->
    </v-container>
</template>
<script>
import ForManager from './comp/ForManager.vue';
import ForNormal from './comp/ForNormal.vue';
import ForBiller from './comp/ForBiller.vue';
export default {
    name:'commanDashboard',
    data(){
        return{

        }
    },
    props:{
        role:String,
    },
    components:{
        ForManager,
        ForNormal,
        ForBiller
    }
}
</script>