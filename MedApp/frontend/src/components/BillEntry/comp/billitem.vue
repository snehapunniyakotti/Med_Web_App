<template>

    <v-container class="pb-2 pb-3 ">
        <v-expansion-panels>
            <v-expansion-panel ref="expansionPanel">
                <v-expansion-panel-header class="white--text font-weight-bold" color="#81D4FA" elevation="0" dense>
                    Item
                </v-expansion-panel-header>
                <v-expansion-panel-content class=" pl-6 " elevation="0">
                    <v-layout class="mt-10">
                        <v-flex>
                            <v-autocomplete :items="medicines" label="Medicine Name" outlined dense
                                v-model="medicinename"></v-autocomplete>

                        </v-flex>
                        <v-flex>
                            <v-text-field v-model.number="qty" label="Quantity" outlined class="pl-2" type="number"
                                dense></v-text-field>
                        </v-flex>
                        <v-flex>
                            <v-btn class="ml-6 white--text" color="#81D4FA" @click="addBillItem">
                                Add +
                            </v-btn>
                        </v-flex>
                    </v-layout>
                </v-expansion-panel-content>
            </v-expansion-panel>
            <v-snackbar v-model="snackbar1" :timeout="timeout" centered color="error">Requested Quantity Unavailable</v-snackbar>
            <v-snackbar v-model="snackbar2" :timeout="timeout" centered color="error"> Incorrect Input, Please Check</v-snackbar>
        </v-expansion-panels>

        <!-- {{ LocalSArr }}<br><br>
        {{ stockarr }}  -->
    </v-container>


</template>

<script>
import EventServices from '../../../services/EventServices';
export default {
    name: "billitem",
    components: {
    },
    mounted() {
        this.fetchMedicines(); // Fetch the list of medicines when component is mounted
    },
    data() {
        return {
            medicinename: '',
            qty: null,
            item: this.$store.state.Stock.map(medicine => medicine.MedicineName),
            msg: '',
            timeout: 2500,
            snackbar1: false,
            snackbar2: false,
            unitprice: 0,
            Amount: 0,
            medarr: [],
            tempmedname: '',
            tempunitprice: 0,
            brand: '',
            amount: 0,
            totalprice: 0,
            CheckingPointArr: [],
            obj: {
                snackbar: false,
                color: 'red',

                topvar: false,
                bottomvar: false,
                text: 'There is No enough quantity available in inventry',
            },
            medicines: [],

        }
    },
    methods: {
        fetchMedicines() {
            EventServices.getMedicineNames()
                .then(response => {
                    const data = response.data;
                    // console.log('API Response:', data); // Debug log
                    if (data.resp.status === "S") {
                        this.medicines = data.medArr.map(med => med.medName); // Ensure correct mapping
                        // console.log('Medicines:', this.medicines); // Debug log
                    } else {
                        console.error('Error:', data.resp.errMsg);
                    }
                })
                .catch(error => {
                    console.error('API Error:', error);
                });
        },

        addBillItem() {
            if (this.medicinename === '' || this.qty <= 0) {
                this.snackbar2 = true;
                // this.medicinename = '';
                this.qty = null;
                return;
            }

            EventServices.checkQuantity(this.medicinename, this.qty)
                .then(response => {
                    const data = response.data;
                    if (data.resp.status === "S") {
                        const billItem = data.oneBillItem;
                        // console.log(billItem);
                        let obj3 = {
                            medicineName: billItem.medicineName,
                            brand: billItem.brand,
                            quantity: this.qty,
                            amount: this.quantity * billItem.unitPrice,
                            unitPrice: billItem.unitPrice,
                        };
                        // console.log(obj3);
                        this.$emit('buyyedpro', obj3);
                        this.medicinename = '';
                        this.qty = null;
                    } else {
                        this.snackbar1 = true;
                        this.medicinename = '';
                        this.qty = null;
                    }
                })
                .catch(error => {
                    console.error('API Error:', error);
                });
        },
        
        brandname() {
            // console.log('a');
            for (let i = 0; i < this.arr.length; i++) {
                if (this.medicinename == this.arr[i].MedicineName) {
                    this.brand = this.arr[i].Brand;
                }
            }
            // console.log(this.brand);
        },
        EachAmount() {
            this.amount = this.qty * this.tempunitprice;
        },
        getLocalArray() {
            for (let i = 0; i < this.stockarr.length; i++) {
                this.LocalSArr.push({
                    MedicineName: this.stockarr[i].MedicineName,
                    quantity: this.stockarr[i].quantity,
                    Unitprice: this.stockarr[i].Unitprice,
                },);
            }

        }
    },

    
    props: {
        roughdata: String,
    },
    watch: {
        roughdata() {
            if (this.roughdata == "true") {
                this.medarr = [];
                this.$emit('changerough');
            }
        }
    },

}
</script>
