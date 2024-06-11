<template>
    <v-container class="mt-12">
        <v-card elevation="4"  >
            <v-card-title>
               <h2>Stock View</h2>
                <v-spacer></v-spacer>
            </v-card-title>
            <v-text-field v-model="search" append-outer-icon="mdi-magnify" label="Search" outlined hide-details
                rounded></v-text-field>
            <v-data-table :headers="headers" :items="stockData" :search="search"  ></v-data-table>
        </v-card>
    </v-container>
</template>
<script>
import EventServices from '../services/EventServices';
export default {
    name: 'stockview',
    data() {
        return {
            search: '',
            stockData: []
        }
    }, 
    methods: {
       
        fetchStockData() {
            EventServices.fetchAllStock()
                .then(response => {
                    const data = response.data;
                    if (data.resp.status === "S") {
                        // console.log(data)
                        this.stockData = data.stockArr;
                    } else {
                        console.error('Error fetching stock data:', data.resp.errMsg);
                    }
                })
                .catch(error => {
                    console.error('API Error:', error);
                });
        }
    },
    computed: {

        headers() {
            return [
                { text: 'Medicice Name', value: 'medicineName' },
                { text: 'Brand', value: 'brand' },
                { text: 'Quantity', value: 'quantity' },
                { text: 'Unit Price', value: 'unitPrice' },
            ]
        },
        
    },
    mounted() {
        this.fetchStockData();
    }
}
</script>
