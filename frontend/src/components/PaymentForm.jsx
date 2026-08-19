import React, { useState } from 'react';
import { initiatePayment } from '../api';

const PaymentForm = () => {
    const [formData, setFormData] = useState({
        amount: '',
        currency: 'ETB',
        description: '',
        customerEmail: '',
        customerName: '',
        phoneNumber: '',
    });
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState('');
    const [paymentUrl, setPaymentUrl] = useState('');

    const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData(prev => ({
            ...prev,
            [name]: value,
        }));
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        setLoading(true);
        setError('');
        setPaymentUrl('');

        try {
            const response = await initiatePayment({
                amount: parseFloat(formData.amount),
                currency: formData.currency,
                description: formData.description,
                customer_email: formData.customerEmail,
                customer_name: formData.customerName,
                phone_number: formData.phoneNumber,
                return_url: window.location.origin + '/payment-success',
                cancel_url: window.location.origin + '/payment-cancel',
            });

            if (response.checkout_url && response.uuid) {
                // Redirect to AddisPay checkout
                window.location.href = `${response.checkout_url}/${response.uuid}`;
            } else {
                setError('Payment initiation failed: No checkout URL returned');
            }
        } catch (err) {
            setError(typeof err === 'string' ? err : 'Payment initiation failed');
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="payment-form-container">
            <h2>Pay with AddisPay</h2>
            {error && <div className="error-message">{error}</div>}
            {paymentUrl && (
                <div className="success-message">
                    Payment initiated! <a href={paymentUrl}>Proceed to payment</a>
                </div>
            )}
            <form onSubmit={handleSubmit}>
                <div className="form-group">
                    <label htmlFor="amount">Amount (ETB)</label>
                    <input
                        type="number"
                        id="amount"
                        name="amount"
                        value={formData.amount}
                        onChange={handleChange}
                        step="0.01"
                        min="0.01"
                        required
                    />
                </div>

                <div className="form-group">
                    <label htmlFor="customerName">Full Name</label>
                    <input
                        type="text"
                        id="customerName"
                        name="customerName"
                        value={formData.customerName}
                        onChange={handleChange}
                        required
                    />
                </div>

                <div className="form-group">
                    <label htmlFor="customerEmail">Email Address</label>
                    <input
                        type="email"
                        id="customerEmail"
                        name="customerEmail"
                        value={formData.customerEmail}
                        onChange={handleChange}
                        required
                    />
                </div>

                <div className="form-group">
                    <label htmlFor="description">Description</label>
                    <input
                        type="text"
                        id="description"
                        name="description"
                        value={formData.description}
                        onChange={handleChange}
                        placeholder="Product/Service description"
                        required
                    />
                </div>

                <div className="form-group">
                    <label htmlFor="phoneNumber">Phone Number</label>
                    <input
                        type="tel"
                        id="phoneNumber"
                        name="phoneNumber"
                        value={formData.phoneNumber}
                        onChange={handleChange}
                        placeholder="2519XXXXXXXX"
                        required
                    />
                </div>

                <button type="submit" disabled={loading}>
                    {loading ? 'Processing...' : 'Pay Now'}
                </button>
            </form>
        </div>
    );
};

export default PaymentForm;