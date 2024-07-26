import 'dart:math';
import 'package:flutter/material.dart';
import 'package:wallet_app/constants.dart';
import '../../../components/appbar/custom_app_bar.dart';
import '../../../components/card/balance_card.dart';
import '../card/card_screen.dart';
import '../history/history_screen.dart';
import '../more/more_screen.dart';
import 'topup/topup_screen.dart';

class HomeScreen extends StatefulWidget {
  @override
  _HomeScreenState createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  int _selectedIndex = 0;
  late double balance = 0;

  void _onItemTapped(int index) {
    setState(() {
      _selectedIndex = index;
    });
  }

  @override
  void initState() {
    super.initState();
    balance = _generateRandomBalance();
  }

  double _generateRandomBalance() {
    final random = Random();
    return random.nextDouble() *
        10000; // Generates a random balance between 0 and 10,000
  }

  void _handleAddMoney() {
    // Implement add money functionality
    print('Add Money pressed');
    // Navigate to Term Condition Screen when tapped
    Navigator.push(
      context,
      MaterialPageRoute(builder: (context) => TopUpScreen()),
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: CustomAppBar(
        logoAssetPath: 'assets/logo.png', // Path to your logo asset
        showBackButton:
            false, // Set this to true if you want to show the back button
        showIconLogout: true,
      ),
      body: _buildBody(),
      bottomNavigationBar: BottomNavigationBar(
        type: BottomNavigationBarType.fixed,
        items: const <BottomNavigationBarItem>[
          BottomNavigationBarItem(
            icon: Icon(Icons.home),
            label: 'Home',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.history),
            label: 'History',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.wallet),
            label: 'Cards',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.manage_accounts_rounded),
            label: 'More',
          ),
        ],
        currentIndex: _selectedIndex,
        selectedItemColor: kPrimaryColor,
        unselectedItemColor: Colors.grey,
        onTap: _onItemTapped,
      ),
    );
  }

  Widget _buildBody() {
    switch (_selectedIndex) {
      case 0:
        return _buildHomeContent();
      case 1:
        return HistoryScreen(); // Navigate to History Screen
      case 2:
        return CardScreen(); // Navigate to Card Screen
      case 3:
        return MoreScreen(); // Navigate to More Screen
      default:
        return Container();
    }
  }

  Widget _buildHomeContent() {
    return Center(
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: <Widget>[
          Padding(
            padding: const EdgeInsets.all(16.0),
            child: BalanceCard(
              balance: balance,
            ),
          ),
          SizedBox(height: 20),
          ElevatedButton(
            onPressed: _handleAddMoney,
            child: Text('Add Money'),
          ),
        ],
      ),
    );
  }
}
