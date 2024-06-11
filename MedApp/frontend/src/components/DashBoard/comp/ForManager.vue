<template>
    <v-container class="mt-12 ">
        <v-card elevation="3" width="800px">
            <v-row>
                <v-col>
                    <v-card-title>TOTAL SALES</v-card-title>
                    <v-card-text>{{ sales }}</v-card-text>
                </v-col>
                <v-col>
                    <v-card-title>CURRENT INVENTRY VALUE</v-card-title>
                    <v-card-text>{{ stockValue }}</v-card-text>
                    <h1></h1>
                </v-col>
            </v-row>
        </v-card>
    </v-container>
</template>
<script>
import EventServices from '../../../services/EventServices';

export default {
    name: 'ForManager',
    data() {
        return {
            sales: 0,
            stockValue: 0,
        }
    },

    methods: {
        fetchManagerDashboard() {
            EventServices.getManagerDashboard()
                .then(response => {
                    const data = response.data;
                    if (data.resp.status === "S") {
                        this.sales = data.arr.totalSales;
                        this.stockValue = data.arr.curInventry;
                    } else {
                        console.error('Error fetching dashboard data:', data.resp.errMsg);
                    }
                })
                .catch(error => {
                    console.error('API Error:', error);
                });
        }
    },
    mounted() {
        this.fetchManagerDashboard();
    }

}
</script>
