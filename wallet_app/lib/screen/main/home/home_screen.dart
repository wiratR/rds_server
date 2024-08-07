import 'dart:math';
import 'package:flutter/material.dart';
import '../../../components/appbar/custom_app_bar.dart';
import '../../../components/card/balance_card.dart';
import '../../../constants.dart';
import '../../../models/user/user_model.dart';
import '../../../service/auth/auth_service.dart';
import '../../../service/user/user_service.dart';
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
  double balance = 0;
  String? token;
  late UserService userService;
  UserDetails? userDetails;
  UserResponse? userResponse;
  late AuthService authService;
  String? userId;

  @override
  void initState() {
    super.initState();
    balance = _generateRandomBalance();
    authService = AuthService();
    WidgetsBinding.instance.addPostFrameCallback((_) async {
      await _initializeServices();
    });
  }

  Future<void> _initializeServices() async {
    token = await authService.getToken();
    if (token != null) {
      userId = await authService.getUserId();
      debugPrint('got a user id = $userId');

      userService = UserService(token!);
      _fetchUserDetails(userId!);
    }
  }

  double _generateRandomBalance() {
    final random = Random();
    return random.nextDouble() *
        10000; // Generates a random balance between 0 and 10,000
  }

  void _handleAddMoney() {
    // Implement add money functionality
    debugPrint('Add Money pressed');
    // Navigate to TopUpScreen when tapped
    Navigator.push(
      context,
      MaterialPageRoute(builder: (context) => TopUpScreen()),
    );
  }

  Future<void> _fetchUserDetails(String userId) async {
    // Implement your logic to fetch user details here
    debugPrint('start _fetchUserDetails');
    userResponse = await userService.getUserById(userId);
    if (userResponse != null) {
      debugPrint('got a userResponse');
      setState(() {
        userDetails = UserDetails.fromUserResponse(userResponse!);
      });
    }
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
        return const HistoryScreen(); // Navigate to History Screen
      case 2:
        return const CardScreen(); // Navigate to Card Screen
      case 3:
        return userDetails != null
            ? MoreScreen(userDetails: userDetails!)
            : const Center(child: CircularProgressIndicator());
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
          const SizedBox(height: 20),
          ElevatedButton(
            onPressed: _handleAddMoney,
            child: const Text('Add Money'),
          ),
        ],
      ),
    );
  }

  void _onItemTapped(int index) {
    setState(() {
      _selectedIndex = index;
    });
    if (index == 3) {
      _fetchUserDetails(userId!);
    }
  }
}
