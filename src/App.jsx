import React from 'react';
import { Route, Routes, BrowserRouter as Router } from 'react-router-dom';
import Home from './pages/home';
import PaymentPage from './pages/PaymentPage';
import SubmitProblemPage from './pages/SubmitProblemPage';
import OpenRentalsPage from './pages/OpenRentalsPage';
import LoginPage from './pages/LoginPage';
import OldPaymentsPage from './pages/OldPaymentsPage';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<LoginPage />} />
        <Route path="/home" element={<Home />} />
        <Route path="/payment" element={<PaymentPage />} />
        <Route path="/submit-problem" element={<SubmitProblemPage />} />
        <Route path="/open-rentals" element={<OpenRentalsPage />} />
        <Route path="/old-payments" element={<OldPaymentsPage />} />
        {/* Add more routes for other pages if needed */}
      </Routes>
    </Router>
  );
}

export default App;
