import React from 'react';
import { Route, Routes, BrowserRouter as Router } from 'react-router-dom';
import ButtonAppBar from './ButtonAppBar'; // Import the ButtonAppBar component
import Home from './pages/Home';
import PaymentPage from './pages/PaymentPage';
import SubmitProblemPage from './pages/SubmitProblemPage';
import RentalManagementPage from './pages/RentalManagementPage';
import OldPaymentsPage from './pages/OldPaymentsPage';
import RegisterPage from './pages/RegisterPage';
import LoginPage from './pages/LoginPage';
import PaymentTest from './pages/PaymentTest';

function App() {
  return (
    <Router>
      <ButtonAppBar /> {/* Include the ButtonAppBar component */}
      <Routes>
        <Route path="/" element={<LoginPage />} />
        <Route path="/home" element={<Home />} />
        <Route path="/payment" element={<PaymentPage />} />
        <Route path="/submit-problem" element={<SubmitProblemPage />} />
        <Route path="/rental-management" element={<RentalManagementPage />} />
        <Route path="/old-payments" element={<OldPaymentsPage />} />
        <Route path="/register" element={<RegisterPage />} />
        <Route path="/payment-test" element={<PaymentTest />} />
        {/* Add more routes for other pages if needed */}
      </Routes>
    </Router>
  );
}

export default App;
