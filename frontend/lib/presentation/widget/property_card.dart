import 'package:flutter/material.dart';
import 'package:intl/intl.dart';
import 'package:phosphor_flutter/phosphor_flutter.dart';


import '../../domain/entity/real_estate.dart';
import '../../l10n/app_localizations.dart';

class PropertyCard extends StatelessWidget {
  final RealEstateEntity property;
  final double quote;

  const PropertyCard({super.key, required this.property, required this.quote});

  @override
  Widget build(BuildContext context) {
    final l10n = AppLocalizations.of(context)!;
    final currencyBrl = NumberFormat.currency(locale: 'pt_BR', symbol: 'R\$');
    final currencyUsd = NumberFormat.currency(locale: 'en_US', symbol: '\$');

    return Card(
      elevation: 0,
      shape: RoundedRectangleBorder(
        side: BorderSide(color: Colors.grey.shade200),
        borderRadius: BorderRadius.circular(12),
      ),
      clipBehavior: Clip.antiAlias,
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Stack(
            children: [

              Container(
                height: 200,
                width: double.infinity,
                color: Colors.grey.shade200,
                child: Center(
                  child: Image.asset(
                    property.imageUrl,
                    height: 200,
                    fit: BoxFit.cover,
                    errorBuilder: (context, error, stackTrace) {
                      return Icon(PhosphorIcons.houseLine(PhosphorIconsStyle.thin), size: 64, color: Colors.grey.shade400);
                    },
                  ),
                ),
              ),
              Positioned(
                top: 12,
                left: 12,
                child: Row(
                  children: [
                    _buildBadge(
                      context,
                      text: property.transactionType == 'SELL' ? l10n.badgeSale : l10n.badgeRent,
                      backgroundColor: const Color(0xFF1A2B47),
                      textColor: Colors.white,
                    ),
                  ],
                ),
              ),
            ],
          ),

          Padding(
            padding: const EdgeInsets.all(16.0),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  property.address.street,
                  style: const TextStyle(fontWeight: FontWeight.bold, fontSize: 16),
                  maxLines: 1,
                  overflow: TextOverflow.ellipsis,
                ),
    const SizedBox(width: 16),
                Row(
                children: [
                Icon(PhosphorIcons.mapPin(), size: 16, color: Colors.grey),
                  const SizedBox(width: 6),
                  Text(
                    property.address.zipCode,
                    style: const TextStyle(fontWeight: FontWeight.bold, fontSize: 16),
                    maxLines: 1,
                    overflow: TextOverflow.ellipsis,
                  ),
                  const SizedBox(width: 6),
                  Text(
                    "${property.address.city} - ${property.address.stateAbbr}",
                    style: const TextStyle(fontWeight: FontWeight.bold, fontSize: 16),
                    maxLines: 1,
                    overflow: TextOverflow.ellipsis,
                  ),

                ],
                ),
                const SizedBox(height: 16),
                Text(
                  currencyBrl.format(property.price),
                  style: const TextStyle(fontWeight: FontWeight.w900, fontSize: 18),
                ),
                Text(
                  currencyUsd.format(property.price * quote),
                  style: TextStyle(color: Colors.grey.shade600, fontSize: 12),
                ),
                const SizedBox(height: 16),
              ],
            ),
          ),
        ],
      ),
    );
  }

  Widget _buildBadge(BuildContext context, {required String text, required Color backgroundColor, required Color textColor}) {
    return Container(
      padding: const EdgeInsets.symmetric(horizontal: 12, vertical: 6),
      decoration: BoxDecoration(
        color: backgroundColor,
        borderRadius: BorderRadius.circular(20),
      ),
      child: Text(
        text,
        style: TextStyle(color: textColor, fontSize: 10, fontWeight: FontWeight.bold),
      ),
    );
  }
}