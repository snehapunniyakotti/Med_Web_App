<template>
    <v-app-bar elevation="10" app class="white--text" color="#81D4FA" dark height="120">
        <v-toolbar-title >Sneha MedApp</v-toolbar-title>
        <v-icon  class="mr-4 ml-4   " size="40">mdi-medical-bag</v-icon>
        <!-- <h1>{{ "You login as " +this.$route.params.role }}</h1> -->
        <v-spacer></v-spacer>
        <v-btn :class="{ 'blue--text white': activeButton === index }" text @click="pathNavigate(item.path, index)"
            v-for="(item, index) in arr" :key="index">{{ item.name }}</v-btn>
    </v-app-bar>
</template>
<script>
import EventServices from "../../services/EventServices";
export default {
    name: 'NavBar',
    data() {
        return {
            arrobj: {
                Biller: [{ name: 'Dashboard', path: 'dashboard' }, { name: 'stock view', path: 'stockview', }, { name: 'bill entry', path: 'billentry' }, { name: 'Log Out', path: '' }],
                Manager: [{ name: 'Dashboard', path: 'dashboard' }, { name: 'stock Entry', path: 'stockentry' }, { name: 'stock view', path: 'stockview' }, { name: 'sales report', path: 'salesreport' }, { name: 'Log Out', path: '' }],
                SystemAdmin: [{ name: 'Dashboard', path: 'dashboard' }, { name: 'Add user', path: 'adduser' }, { name: 'login history', path: 'loginhistory' }, { name: 'Log Out', path: '' }],
                Inventry: [{ name: 'Dashboard', path: 'dashboard' }, { name: 'stock Entry', path: 'stockentry' }, { name: 'stock view', path: 'stockview' }, { name: 'Log Out', path: '' } ],
                // , { name: 'samp', path: 'samp' }
            },
            user_type: '',
            user_id: '',
            user_HistoryId: '',
            activeButton: 0,
           
        }
    },
    props: {
        type: String,
        id: String,
        lid: Number,
    },
    computed: {
        arr() {
            return this.arrobj[this.type] || []
        }
    },
    watch: {
        '$route'() {
            // console.log(to,from)
            this.setActiveButton(); // Update active button on route change
        }
    },
    mounted() {
        this.setActiveButton(); // Set the active button on component mount
    },
    methods: {
        pathNavigate(path, index) {
            if (this.$route.path == `/${path}`) {
                return;
            }
            if (path == '') {
                this.logout();
                return;
            }
            // console.log(path + ' ' + this.type + ' ' + this.id + ' ' +this.lid );
            this.user_type = this.type
            this.user_id = this.id
            this.user_HistoryId = this.lid
            this.$router.push({ name: path, params: { role: this.user_type, Id: this.user_id, lId: this.user_HistoryId } })
            this.activeButton = index;
        },
        logout() {
            let user = this.$route.params.Id;
            if (!user) {
                this.$router.push('/Errormsg')
            }
            // console.log(this.lid)
            EventServices.updateLoginHistory(this.lid)
                .then(response => {
                    if (response.data && response.data.resp.status === 'S') {
                        const date = new Date();
                        let obj = {
                            User_Id: user,
                            Type: 'Logout',
                            Date: date,
                        };
                        this.$store.commit('login_entry', obj);
                        this.$router.push('/');
                    } else {
                        console.error('Logout failed:', response.data.resp.errMsg);
                    }
                })
                .catch(error => {
                    console.error('API Error:', error);
                    this.$router.push('/Errormsg');
                });
        },
        setActiveButton() {
            const currentPath = this.$route.path.replace('/', '');
            const activeIndex = this.arr.findIndex(item => item.path === currentPath);
            this.activeButton = activeIndex !== -1 ? activeIndex : 0;
        }

    }

}
</script>
