import 'package:frontend/data/datasources/exchange_rate.dart';

import '../../domain/entity/exchange_rate.dart';
import '../../domain/repository/exchange_rate.dart';

class ExchangeRateRepositoryImpl implements ExchangeRateRepository{
  final ExchangeRateDataSource exchangeRateDataSource;

  ExchangeRateRepositoryImpl({required this.exchangeRateDataSource});

  @override
  Future<ExchangeRateEntity> getLatestExchangeRate() async {
    final res =  await exchangeRateDataSource.getLatestExchangeRate();

    return res;
  }

  @override
  Future<String> addExchangeRate(ExchangeRateEntity exchangeRate) async {
    final res =  await exchangeRateDataSource.addExchangeRate(exchangeRate);

    return res;
  }

}

