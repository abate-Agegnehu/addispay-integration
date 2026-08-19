import React, { useState, useEffect } from 'react';
import { useSearchParams } from 'react-router-dom';
import { getPaymentStatus } from '../api';

const PaymentStatus = () => {
  const [searchParams] = useSearchParams();
  const [status, setStatus] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  const transactionId = searchParams.get('transaction_id');
  const statusParam = searchParams.get('status');

  useEffect(() => {
    const checkStatus = async () => {
      if (!transactionId) {
        setLoading(false);
        setError('No transaction ID provided');
        return;
      }

      try {
        const data = await getPaymentStatus(transactionId);
        setStatus(data);
      } catch (err) {
        setError(typeof err === 'string' ? err : 'Failed to get payment status');
      } finally {
        setLoading(false);
      }
    };

    // If we have a status from redirect, show it immediately
    if (statusParam) {
      setStatus({
        status: statusParam,
        transactionId: transactionId,
      });
      setLoading(false);
    } else {
      checkStatus();
    }
  }, [transactionId, statusParam]);

  if (loading) {
    return <div className="loading">Checking payment status...</div>;
  }

  if (error) {
    return <div className="error">{error}</div>;
  }

  if (!status) {
    return <div className="no-status">No payment status available</div>;
  }

  const getStatusIcon = () => {
    switch (status.status) {
      case 'completed':
        return '✅';
      case 'pending':
        return '⏳';
      case 'failed':
        return '❌';
      case 'cancelled':
        return '🚫';
      default:
        return '❓';
    }
  };

  const getStatusMessage = () => {
    switch (status.status) {
      case 'completed':
        return 'Payment completed successfully!';
      case 'pending':
        return 'Payment is pending. We\'ll notify you when it\'s complete.';
      case 'failed':
        return 'Payment failed. Please try again.';
      case 'cancelled':
        return 'Payment was cancelled.';
      default:
        return 'Unknown payment status';
    }
  };

  return (
    <div className="payment-status-container">
      <h2>Payment Status</h2>
      <div className={`status-card ${status.status}`}>
        <div className="status-icon">{getStatusIcon()}</div>
        <div className="status-message">{getStatusMessage()}</div>
        {status.transactionId && (
          <div className="transaction-id">
            Transaction ID: {status.transactionId}
          </div>
        )}
        {status.amount && (
          <div className="amount">
            Amount: {status.amount} {status.currency || 'ETB'}
          </div>
        )}
      </div>
    </div>
  );
};

export default PaymentStatus;