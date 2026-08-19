import React from 'react';
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import PaymentForm from './components/PaymentForm';
import PaymentStatus from './components/PaymentStatus';
import PaymentSuccess from './components/PaymentSuccess';
import './App.css';

function App() {
    return (
        <Router>
            <div className="App">
                <nav>
                    <Link to="/">Home</Link>
                </nav>
                <main>
                    <Routes>
                        <Route path="/" element={
                            <div className="home">
                                <h1>AddisPay Payment Integration</h1>
                                <PaymentForm />
                            </div>
                        } />
                        <Route path="/payment-success" element={<PaymentSuccess />} />
                        <Route path="/payment-cancel" element={<PaymentStatus />} />
                    </Routes>
                </main>
            </div>
        </Router>
    );
}

export default App;