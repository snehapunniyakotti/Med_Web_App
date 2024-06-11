<template>
    <div class="text-center " >
        <v-dialog v-model="dialog" width="500px">
            <template v-slot:activator="{ on, attrs }">
                <v-btn color="#81D4FA" class="white--text" v-bind="attrs" v-on="on">
                    Preview
                </v-btn>
            </template>
            <v-card>
                <v-card-text outline >
                    <v-layout>
                        <v-flex md-12 text-center class="text-h5 mt-3 black--text text--lighten-4 font-weight-bold "> Sneha Medicals
                            <v-icon size="30">mdi-truck-plus-outline</v-icon>
                        </v-flex>
                        
                    </v-layout>
                    <v-layout class="ma-3 lighten-1 rounded-lg light-blue lighten-3"  >
                        <v-flex class="ma-3 white--text text-subtitle-1 font-weight-bold">Medicine name</v-flex>
                        <v-flex class="ma-3 white--text text-subtitle-1 font-weight-bold">Qty</v-flex>
                        <v-flex class="ma-3 white--text text-subtitle-1 font-weight-bold">Amount</v-flex>
                    </v-layout>
                    <v-layout v-for="val,idx in Previewarr" :key="idx" class="pl-2 pb-3">
                        <v-flex lg4 text-center>{{ val.medicineName }}</v-flex>
                        <v-flex lg4 text-center>{{ val.quantity }}</v-flex>
                        <v-flex lg4 text-center>{{ val.amount }}</v-flex>
                    </v-layout>
                    <v-layout class="mb-3">
                        <v-flex lg4></v-flex>
                        <v-flex lg4 text-right class="black--text pr-5 subtitle-1"><b>Total</b></v-flex>
                        <v-flex lg4 text-center class="black--text font-weight-medium ">{{ Total }}</v-flex>
                    </v-layout>
                    <v-layout class="mb-3">
                        <v-flex lg4></v-flex>
                        <v-flex lg4 text-right class="black--text pr-5 subtitle-1"><b>GST</b></v-flex>
                        <v-flex lg4 text-center class="black--text font-weight-medium ">{{ GST }}</v-flex>
                    </v-layout>

                    <v-layout class="mb-3">
                        <v-flex lg4></v-flex>
                        <v-flex lg4 text-right class="black--text font-weight-black pr-5 subtitle-1"><b>Net Price</b></v-flex>
                        <v-flex lg4 text-center class="black--text font-weight-medium ">{{ NetPrice }}</v-flex>
                    </v-layout>

                </v-card-text>
                <v-card-actions>
                    <v-btn dark color="#81D4FA" class="mb-3" @click="dialog = false">Close <v-icon> mdi-close-circle</v-icon></v-btn>
                    <v-spacer></v-spacer>
                    <!-- <v-btn dark color="#81D4FA" class="mb-3" @click="printto()">Print</v-btn> -->
                </v-card-actions>
            </v-card>
        </v-dialog>
    </div>
</template>
<script>
export default {
    name: '',
    data() {
        return {
            arr: ['paracetomol', 5, 250],
            billMaster: this.$store.state.BillMaster,
            billDetails: this.$store.state.BillDetails,
            dialog: false,

        }
    },
    props:{
        Previewarr:Array,
        Total:Number, 
        GST:Number,
        NetPrice:Number, 
    },
    methods:{
        printto(){ 
          this.dialog = false;

        }
    }

}
</script>