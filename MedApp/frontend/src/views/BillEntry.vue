<template>
  <v-container class="mt-16">
    <v-layout></v-layout>
    <v-layout class="mt-5">
      <v-flex>
        <!-- @changerough="roughemit1" -->
        <billitem @buyyedpro="amount" :roughdata="rough" />

        <billtable :billNumber="billno" :userArr="Medarr" :tp="totalprice" @billnochange="billchange"
          @emptypreview="emptyit" @billclosed="billclosed" />

        <v-snackbar v-model="snackbar3" :timeout="timeout" centered color="error"> limited stocks left</v-snackbar>
      </v-flex>
    </v-layout>


  </v-container>
</template>

<script>
import billitem from '../components/BillEntry/comp/billitem.vue';
import billtable from '../components/BillEntry/comp/billtable.vue';
import EventServices from '../services/EventServices';


export default {
  name: "BillEntry",
  components: {
    billitem, billtable,
  },
  // computed: {
  //   billno() {
  //     return `${Math.floor(Math.random() * 9000) + 1000+ 1}`;
  //   },
  // },
  data() {
    return {
      Medarr: [],
      totalprice: 0,
      rough: "",
      timeout: 2500,
      snackbar3: false,
      billno: `${Math.floor(Math.random() * 9000) + 1000}`,
      
    }
  },
  
  methods: {
    billclosed() {
      this.Medarr = [];
      this.totalprice = 0;
    },
    billchange() {
      // this.billno += 1;
      this.billno = `${parseInt(this.billNumber) + 1}`;
    },
    emptyit() {
      this.Medarr = [];
      this.totalprice = 0;
      this.rough = "true";
    },
    roughemit1() {
      this.rough = "";
    },

    amount(obj) {
      // Fetch the original stock quantity from the API
      EventServices.getMedicineDetails(obj.medicineName)
        .then(response => {
          const data = response.data;
          if (data.resp.status === "S" && data.medDetails) { // Check if medDetails is defined
            const originalStock = data.medDetails;

            if (originalStock.quantity !== undefined) { // Check if Quantity is defined
              let oriqnt = originalStock.quantity;
              let oldobj = this.Medarr.find(med => med.medicineName === obj.medicineName);

              if (oldobj !== undefined) {

                if (oriqnt < (oldobj.quantity + obj.quantity)) {
                  this.snackbar3 = true;
                  return;
                }
                oldobj.quantity += obj.quantity;
                oldobj.amount = (oldobj.quantity * obj.unitPrice);
              } else {

                if (oriqnt < obj.quantity) {
                  this.snackbar3 = true;
                  return;
                }
                obj.amount = obj.unitPrice * obj.quantity

                this.Medarr.push(obj);
              // console.log(this.Medarr)

              }
              // Update Medarr with unique items
              this.Medarr = this.Medarr.reduce((acc, cur) => {
                const existing = acc.find(item => item.medicineName === cur.medicineName);
                if (existing) {
                  existing.quantity += cur.quantity;
                  existing.amount = existing.quantity * existing.unitPrice;
                } else {
                  acc.push(cur);
                }
                return acc;
              }, []);
              // console.log(this.Medarr)


              // Update totalprice
              this.totalprice = this.Medarr.reduce((total, item) => {
                const itemAmount = parseFloat(item.amount);
                return isNaN(itemAmount) ? total : total + itemAmount;
              }, 0);
              this.totalprice = parseFloat(this.totalprice.toFixed(2));

              
            } else {
              console.error('Error: Quantity is undefined'); // Added error log for undefined Quantity
            }
          } else {
            console.error('Error:', data.resp.errMsg || 'No medDetails found'); // Updated error log
          }
        })
        .catch(error => {
          console.error('API Error:', error);
        });
    }



  },


};
</script>