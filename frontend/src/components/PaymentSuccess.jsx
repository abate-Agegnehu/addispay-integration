import React from 'react';
import { Link, useSearchParams } from 'react-router-dom';

const PaymentSuccess = () => {
    const [searchParams] = useSearchParams();
    const uuid = searchParams.get('uuid');
    const transactionReference = searchParams.get('tx_ref') || searchParams.get('transaction_id');
    const amount = searchParams.get('total_amount') || searchParams.get('amount');
    const currency = searchParams.get('currency') || 'ETB';

    return (
        <section className="payment-result success-result" aria-labelledby="success-title">
            <div className="result-mark" aria-hidden="true">&#10003;</div>
            <p className="result-eyebrow">AddisPay confirmation</p>
            <h1 id="success-title">Payment successful</h1>
            <p className="result-intro">
                Your payment has been received. Keep the details below for your records.
            </p>

            {(amount || transactionReference || uuid) && (
                <dl className="payment-details">
                    {amount && (
                        <div>
                            <dt>Amount</dt>
                            <dd>{amount} {currency}</dd>
                        </div>
                    )}
                    {transactionReference && (
                        <div>
                            <dt>Reference</dt>
                            <dd>{transactionReference}</dd>
                        </div>
                    )}
                    {uuid && (
                        <div>
                            <dt>Payment ID</dt>
                            <dd>{uuid}</dd>
                        </div>
                    )}
                </dl>
            )}

            <Link className="result-action" to="/">Make another payment</Link>
        </section>
    );
};

export default PaymentSuccess;