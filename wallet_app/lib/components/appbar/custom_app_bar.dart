import 'package:flutter/material.dart';

class CustomAppBar extends StatelessWidget implements PreferredSizeWidget {
  final String logoAssetPath;
  final bool showBackButton;
  final bool showIconLogout;

  CustomAppBar({
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
                  icon: Icon(Icons.arrow_back),
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
                icon: Icon(Icons.logout),
                onPressed: () {
                  // Add your logout functionality here
                },
              ),
            ]
          : null,
    );
  }

  @override
  Size get preferredSize => Size.fromHeight(kToolbarHeight);
}
