import 'package:flutter/material.dart';
import 'package:phosphor_flutter/phosphor_flutter.dart';
import '../../l10n/app_localizations.dart';

class Header extends StatelessWidget {
  final Function(String page) navigate;

  const Header({super.key, required this.navigate});


  @override
  Widget build(BuildContext context) {
    final l10n = AppLocalizations.of(context)!;
    final darkBlue = const Color(0xFF1A2B47);

    return Row(
      mainAxisAlignment: MainAxisAlignment.spaceBetween,
      children: [
        Row(
          children: [
            Icon(PhosphorIcons.houseLine(), color: darkBlue),
            const SizedBox(width: 8),
            Text(
              l10n.appTitle,
              style: TextStyle(fontSize: 20, fontWeight: FontWeight.bold, color: darkBlue),
            ),
          ],
        ),
        Row(
          children: [
            TextButton.icon(
              onPressed: () => navigate("quote"),
              icon: Icon(PhosphorIcons.currencyCircleDollar(), size: 18, color: darkBlue),
              label: Text(l10n.quoteAction, style: TextStyle(color: darkBlue)),
            ),
            const SizedBox(width: 16),
            ElevatedButton.icon(
              onPressed: () =>  navigate("add"),
              style: ElevatedButton.styleFrom(
                  backgroundColor: darkBlue,
                  foregroundColor: Colors.white,
                  shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(8)),
                  padding: const EdgeInsets.symmetric(horizontal: 20, vertical: 16)
              ),
              icon: const Icon(Icons.add, size: 18),
              label: Text(l10n.newAdButton),
            ),
          ],
        ),
      ],
    );
  }
}