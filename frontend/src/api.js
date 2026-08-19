import axios from 'axios';

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080';

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

export const initiatePayment = async (paymentData) => {
  try {
    const response = await api.post('/api/payments/initiate', paymentData);
    return response.data;
  } catch (error) {
    console.error('Payment initiation error:', error);
    throw error.response?.data || error.message;
  }
};

export const getPaymentStatus = async (transactionId) => {
  try {
    const response = await api.get(`/api/payments/status?transaction_id=${transactionId}`);
    return response.data;
  } catch (error) {
    console.error('Payment status check error:', error);
    throw error.response?.data || error.message;
  }
};

export default api;