
import 'package:flutter/material.dart';

import '../../l10n/app_localizations.dart';

class Pagination extends StatelessWidget {
  final int currentPage;
  final int totalPages;

  final Function(int newPage) onPageChanged;

  const Pagination({
    super.key,
    required this.currentPage,
    required this.totalPages,
    required this.onPageChanged,
  });

  @override
  Widget build(BuildContext context) {
    final l10n = AppLocalizations.of(context)!;

    return Padding(
      padding: const EdgeInsets.symmetric(vertical: 16.0),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          ElevatedButton.icon(
            onPressed: currentPage > 1
                ? () => onPageChanged(currentPage - 1)
                : null,
            icon: const Icon(Icons.chevron_left),
            label: Text(l10n.prevPage),
          ),

          const SizedBox(width: 24),

          Text(
            '${l10n.page} $currentPage de $totalPages',
            style: const TextStyle(fontWeight: FontWeight.bold, fontSize: 16),
          ),

          const SizedBox(width: 24),

          ElevatedButton.icon(
            onPressed: currentPage < totalPages && totalPages > 0
                ? () => onPageChanged(currentPage + 1)
                : null,
            icon: const Icon(Icons.chevron_right),
            label: Text(l10n.nextPage),
            iconAlignment: IconAlignment.end,
          ),
        ],
      ),
    );
  }
}