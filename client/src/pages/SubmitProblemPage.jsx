import React, { useState } from 'react';
import { Container, Typography, TextField, Button } from '@mui/material';


const SubmitProblemPage = () => {
  const [problem, setProblem] = useState('');

  const handleSubmitProblem = () => {
    // Logic to handle submitting the problem, e.g., sending it to a server
    console.log(`Submitted problem: ${problem}`);
    // Redirect or show success message after submission
  };

  const handleProblemChange = (event) => {
    setProblem(event.target.value);
  };

  return (
    <Container maxWidth="md" sx={{ textAlign: 'center', mt: 4 }}>
      <Typography variant="h3" gutterBottom>
        Submit a Problem
      </Typography>
      <Typography variant="body1" gutterBottom>
        Describe the problem you are experiencing:
      </Typography>

      <TextField
        label="Problem Description"
        variant="outlined"
        multiline
        rows={4}
        value={problem}
        onChange={handleProblemChange}
        sx={{ mt: 2, width: '100%' }}
      />

      <Button
        variant="contained"
        color="primary"
        onClick={handleSubmitProblem}
        disabled={!problem}
        sx={{ mt: 2 }}
      >
        Submit
      </Button>

      {/* Additional UI for feedback or confirmation of problem submission */}
    </Container>
  );
};

export default SubmitProblemPage;
