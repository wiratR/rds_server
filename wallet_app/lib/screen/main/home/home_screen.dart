// ignore_for_file: use_build_context_synchronously

import 'dart:math';
import 'package:flutter/material.dart';
import '../../../components/appbar/custom_app_bar.dart';
import '../../../components/card/balance_card.dart';
import '../../../constants.dart';
import '../../../models/user/user_model.dart';
import '../../../service/auth/auth_service.dart';
import '../../../service/user/user_service.dart';
import '../../../service/account/account_service.dart';
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
  bool showBalance = true; // Flag to control balance visibility

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
    try {
      token = await authService.getToken();
      if (token != null) {
        userId = await authService.getUserId();
        if (userId != null) {
          debugPrint('got a user id = $userId');
          userService = UserService(token!);
          await _fetchUserDetails(userId!);
        } else {
          debugPrint('User ID is null');
        }
      } else {
        debugPrint('Token is null');
      }
    } catch (e) {
      debugPrint('Error initializing services: $e');
    }
  }

  Future<void> _fetchUserDetails(String userId) async {
    debugPrint('start _fetchUserDetails');
    try {
      userResponse = await userService.getUserById(userId);
      if (userResponse != null) {
        debugPrint('got a userResponse');
        debugPrint(userResponse?.accountId);

        if (userResponse?.accountId == '00000000-0000-0000-0000-000000000000') {
          setState(() {
            showBalance = false; // Do not show balance
            userDetails = UserDetails.fromUserResponse(userResponse!);
          });
        } else {
          setState(() {
            showBalance = true; // Show balance
            userDetails = UserDetails.fromUserResponse(userResponse!);
          });
        }
      }
    } catch (e) {
      debugPrint('Error fetching user details: $e');
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

  // void _handleBindAccount() {
  //   // Implement bind account functionality
  //   debugPrint('Bind Account pressed');
  //   // Navigate to BindAccountScreen when tapped
  //   Navigator.push(
  //     context,
  //     MaterialPageRoute(builder: (context) => BindAccountScreen()),
  //   );
  // }

  void _handleBindAccount() async {
    debugPrint('Bind Account pressed');

    if (userId != null && token != null) {
      // Call the account creation service
      AccountService accountService = AccountService(token!);
      String? response = await accountService.createAccount('mobile', userId!);

      if (response != null) {
        debugPrint('Account created successfully: $response');

        // Refresh the home page by calling _fetchUserDetails again
        await _fetchUserDetails(userId!);

        // Optionally, show a success message
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(content: Text('Account created successfully!')),
        );
      } else {
        debugPrint('Failed to create account');
        // Show an error message
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(content: Text('Failed to create account.')),
        );
      }
    } else {
      debugPrint('User ID or token is null');
      // Handle error cases where userId or token is missing
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
          if (showBalance) ...[
            Padding(
              padding: const EdgeInsets.all(10.0),
              child: BalanceCard(
                balance: balance,
              ),
            ),
            const SizedBox(height: 20),
            ElevatedButton(
              onPressed: _handleAddMoney,
              child: const Text('Add Money'),
            ),
          ] else
            ElevatedButton(
              onPressed: _handleBindAccount,
              child: const Text('Bind Account'),
            ),
        ],
      ),
    );
  }

  void _onItemTapped(int index) {
    setState(() {
      _selectedIndex = index;
    });
    if (index == 3 && userId != null) {
      _fetchUserDetails(userId!);
    } else {
      debugPrint('User ID is null or not set');
    }
  }
}
