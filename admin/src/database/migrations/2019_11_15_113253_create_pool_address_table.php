<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class CreatePoolAddressTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('pool_address', function (Blueprint $table) {
            $table->increments('id');
            $table->string('address', "100")->index()->comment('地址');
            $table->string('pool_id')->index()->comment('矿池id');
            $table->string('coin_id')->index()->comment('币种id');
            $table->tinyInteger('type')->index()->comment('类型');
            $table->integer('created_at_ts')->comment('创建时间戳');
            $table->integer('updated_at_ts')->comment('更新时间戳');
            $table->timestamps();
            $table->unique(["coin_id", "pool_id", "type", "address"]);
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('pool_address');
    }
}
