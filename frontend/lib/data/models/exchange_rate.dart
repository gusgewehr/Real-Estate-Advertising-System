
import 'package:frontend/domain/entity/exchange_rate.dart';

class ExchangeRateModel extends ExchangeRateEntity {
  ExchangeRateModel({
    required super.value,
  });


  factory ExchangeRateModel.fromJSON(Map<String, dynamic> json) {
    return ExchangeRateModel(
      value: json['value'] ?? 0.0,
    );
  }


  Map<String, dynamic> toJson() {
    return {
      'value': super.value,
    };

  }

  factory ExchangeRateModel.fromEntity(ExchangeRateEntity entity) {
    return ExchangeRateModel(
        value: entity.value
    );
  }
}