import 'package:flutter/material.dart';
import 'package:phosphor_flutter/phosphor_flutter.dart';

import '../../l10n/app_localizations.dart';

class CustomError extends StatelessWidget {
  const CustomError({super.key});

  @override
  Widget build(BuildContext context) {
    final l10n = AppLocalizations.of(context)!;



    return Container(
      width: double.infinity,
      height: double.infinity,
      decoration: BoxDecoration(
        // Um degradê suave para não ficar um cinzento "morto"
        gradient: LinearGradient(
          begin: Alignment.topLeft,
          end: Alignment.bottomRight,
          colors: [
            Colors.grey.shade100,
            Colors.grey.shade300,
          ],
        ),
      ),
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          Icon(
            PhosphorIcons.warning(PhosphorIconsStyle.bold),
            size: 48,
            color: Colors.grey.shade500,
          ),
          const SizedBox(height: 8),
          Text(
            l10n.genericError,
            style: TextStyle(
              color: Colors.grey.shade600,
              fontSize: 12,
              fontWeight: FontWeight.w500,
            ),
          ),
        ],
      ),
    );
  }
}