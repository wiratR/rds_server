import 'package:flutter/material.dart';

import '../../screen/login/login_screen.dart';
import '../../service/auth/auth_service.dart';

class CustomAppBar extends StatelessWidget implements PreferredSizeWidget {
  final String logoAssetPath;
  final bool showBackButton;
  final bool showIconLogout;
  final AuthService authService = AuthService();

  CustomAppBar({
    super.key,
    required this.logoAssetPath,
    this.showBackButton = false,
    this.showIconLogout = false,
  });

  @override
  Widget build(BuildContext context) {
    return AppBar(
      leading: showBackButton
          ? Row(
              mainAxisSize: MainAxisSize.min,
              children: [
                IconButton(
                  icon: const Icon(Icons.arrow_back),
                  onPressed: () {
                    Navigator.of(context).pop();
                  },
                ),
                Flexible(
                  child: GestureDetector(
                    onTap: () {
                      Navigator.of(context).pop();
                    },
                  ),
                ),
              ],
            )
          : null,
      title: Row(
        mainAxisAlignment: MainAxisAlignment.center,
        mainAxisSize: MainAxisSize.min,
        children: [
          Image.asset(
            logoAssetPath,
            height: 40, // Adjust the height as needed
          ),
        ],
      ),
      centerTitle: true,
      automaticallyImplyLeading:
          false, // Prevents default back button when showBackButton is false
      actions: showIconLogout
          ? [
              IconButton(
                icon: const Icon(Icons.logout),
                onPressed: () async {
                  // Add your logout functionality here
                  await authService.logout();
                  Navigator.pushReplacement(
                    // ignore: use_build_context_synchronously
                    context,
                    MaterialPageRoute(builder: (context) => LoginScreen()),
                  );
                },
              ),
            ]
          : null,
    );
  }

  @override
  Size get preferredSize => const Size.fromHeight(kToolbarHeight);
}
