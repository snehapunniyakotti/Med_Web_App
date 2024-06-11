import axios from 'axios';

// Create a custom Axios instance with base URL and other configurations
const eventApi = axios.create({
    baseURL: 'http://localhost:8081/', // url of backend running port
    withCredentials: false,
    headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
    }
});

const EventServices = {
    login(userId, password) {
        return eventApi.post('/login', { userId: userId, password: password });
    },
    updateLoginHistory(loginHistoryId) {
        return eventApi.put('/updateloginHistory', { loginHistoryId: loginHistoryId });
    },
    addStock(medicineName, brand) {
        return eventApi.post('/addStock', { medicineName: medicineName, brand: brand });
    },
    getMedicineNames() {
        return eventApi.get('/getMedicineNames');
    },
    getMedicineDetails(medicineName) {
        return eventApi.get(`/mndetails?medicineName=${medicineName}`)
    },
    updateStock(stockData) {
        return eventApi.put('/updatestock', stockData);
    },
    fetchAllStock() {
        return eventApi.get('/getStockView');
    },
    getManagerDashboard() {
        return eventApi.get('/managerDashboard');
    },
    getHistory() {
        return eventApi.get('/history');
    },
    postUser(userId, password, role) {
        return eventApi.post('/addUser', { userId, password, role });
    },
    checkQuantity(medicineName, quantity) {
        return eventApi.post('/checkQuantity', { medicineName, quantity });
    },
    saveBill(billData) {
        return eventApi.post('/saveBill', billData);
    },
    getSalesReport(fromDate, toDate) {
        // Use Axios's params property to pass query parameters
        return eventApi.get('/salesReport', { params: { fromDate: fromDate, toDate: toDate }});
    },
    getBillerDashboard(userId) {
        return eventApi.post('/billerDashboard',userId);
    },

};

export default EventServices;