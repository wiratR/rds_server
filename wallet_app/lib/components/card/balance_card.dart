import 'package:flutter/material.dart';
import 'package:wallet_app/constants.dart';

class BalanceCard extends StatelessWidget {
  // final String cardHolder;
  // final String cardNumber;
  // final String expiryDate;
  final double balance;

  BalanceCard({
    // required this.cardHolder,
    // required this.cardNumber,
    // required this.expiryDate,
    required this.balance,
  });

  @override
  Widget build(BuildContext context) {
    return Card(
      elevation: 5,
      shape: RoundedRectangleBorder(
        borderRadius: BorderRadius.circular(15),
      ),
      child: Container(
        padding: EdgeInsets.all(50),
        decoration: BoxDecoration(
          gradient: LinearGradient(
            colors: [kPrimaryColor, kPrimaryColor],
            begin: Alignment.topLeft,
            end: Alignment.centerLeft,
          ),
          borderRadius: BorderRadius.circular(15),
        ),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(
              'Balance',
              style: TextStyle(
                color: Colors.white,
                fontSize: 18,
              ),
            ),
            SizedBox(height: 10),
            Text(
              '\$${balance.toStringAsFixed(2)}',
              style: TextStyle(
                color: Colors.white,
                fontSize: 32,
                fontWeight: FontWeight.bold,
              ),
            ),
            SizedBox(height: 20),
          ],
        ),
      ),
    );
  }
}
